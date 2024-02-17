package services

import (
	"context"
	"fmt"

	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type accountService struct {
	store             postgres.UnitOfWork
	accounts          map[string]chan *models.AccountAggregate
	blockchainService ports.BlockchainService
}

func NewAccountService(blockchainService ports.BlockchainService, store postgres.UnitOfWork) ports.AccountService {
	service := &accountService{
		store: store,
	}

	var chains []string
	if err := viper.UnmarshalKey("available_chains", &chains); err != nil {
		log.Fatal().Err(err).Msgf("unable to read available chains from config, %v", err)
	}

	accounts := map[string]chan *models.AccountAggregate{}
	for _, chain := range chains {
		relayerCount := viper.GetInt(fmt.Sprintf("%s.relayer_count", chain))
		accounts[chain] = make(chan *models.AccountAggregate, relayerCount)
	}

	return service
}

func (service *accountService) CreateAccounts(ctx context.Context, chain string, accountCount int32) ([]*models.AccountAggregate, error) {
	accounts := make([]*models.AccountAggregate, 0, accountCount)
	for len(accounts) != int(accountCount) {
		newAccount, err := models.NewAccount()
		if err != nil {
			log.Warn().Msgf("failed to create a new account for %s", chain)
		}
		accounts = append(accounts, models.NewAccountAggregate(*newAccount, 1))
	}

	accountsToSave := lo.Map(accounts, func(item *models.AccountAggregate, _ int) models.Account {
		return item.Account
	})
	if err := service.store.Repo().Account().SaveMany(ctx, &accountsToSave); err != nil {
		log.Error().Err(err)
		return nil, pkg.ErrInternal
	}

	return accounts, nil
}

func (service *accountService) GetAccountForChain(ctx context.Context, chain string) {

}

func (service *accountService) fundAccounts() {

}
