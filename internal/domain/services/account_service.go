package services

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/models/account"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type accountService struct {
	allAccounts       []accounts.Account
	accountQueue      account.ChainToAccountQueue
	blockchainService ports.BlockchainService
	store             postgres.UnitOfWork
}

func NewAccountService(blockchainService ports.BlockchainService, store postgres.UnitOfWork) ports.AccountService {
	service := &accountService{
		accountQueue:      make(account.ChainToAccountQueue),
		blockchainService: nil,
		store:             store,
	}

	var chains []string
	if err := viper.UnmarshalKey("available_chains", &chains); err != nil {
		log.Fatal().Err(err).Msgf("unable to read available chains from config, %v", err)
	}

	accounts := map[string]chan *account.AccountAggregate{}
	for _, chain := range chains {
		relayerCount := viper.GetInt(fmt.Sprintf("%s.relayer_count", chain))
		accounts[chain] = make(chan *account.AccountAggregate, relayerCount)
	}

	return service
}

func (service *accountService) initializeAccountsForChain(ctx context.Context, chain string, accountAmount int32) error {
	accounts, err := service.getAccountsAndMarkAsBeingInUse(ctx, chain, accountAmount)
	if err != nil {
		log.Error().Err(err)
		return pkg.ErrInternal
	}

	usersToNoncesMap, err := service.blockchainService.GetNoncesForAccounts(ctx, accounts)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	if len(accounts) != int(accountAmount) {

	}

}

func (service *accountService) getAccountsAndMarkAsBeingInUse(ctx context.Context, chain string, accountAmount int32) ([]account.Account, error) {
	var accounts []account.Account
	err := service.store.RunInTx(ctx, func(store postgres.UnitOfWorkStore) error {
		accs, err := store.Account().GetAvailableAccountsForChain(ctx, chain, accountAmount)
		if err != nil {
			log.Error().Err(err).Msgf("failed to get accounts for %v", chain)
			return pkg.ErrInternal
		}

		for i := 0; i < len(accs); i++ {
			accs[i].MarkAsBeingUsed()
		}

		updateCount, err := store.Account().UpdateMany(ctx, &accs)
		if err != nil {
			log.Error().Err(err).Msgf("failed to mark accounts as being in use for %s", chain)
			return pkg.ErrInternal
		}
		if int(updateCount) != len(accs) {
			log.Error().Err(err).Msgf("failed to mark accounts as being in use for %s", chain)
			return pkg.ErrInternal
		}
		accounts = accs
		return nil
	})

	return accounts, err
}

func (service *accountService) createAccounts(ctx context.Context, chain string, accountCount int32) ([]account.Account, error) {
	accounts := make([]account.Account, 0, accountCount)
	for len(accounts) != int(accountCount) {
		newAccount, err := account.NewAccount(chain, true)
		if err != nil {
			log.Warn().Msgf("failed to create a new account for %s", chain)
		}
		accounts = append(accounts, *newAccount)
	}

	if err := service.store.Repo().Account().SaveMany(ctx, &accounts); err != nil {
		log.Error().Err(err)
		return nil, pkg.ErrInternal
	}

	return accounts, nil
}

func (service *accountService) GetAccountForChain(ctx context.Context, chain string) {

}

func (service *accountService) fundAccounts() {

}
