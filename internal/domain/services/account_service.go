package services

import (
	"context"
	"fmt"
	"log"

	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/internal/domain/ports"
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
		log.Fatalf("unable to read available chains from config, %v", err)
	}

	accounts := map[string]chan *models.AccountAggregate{}
	for _, chain := range chains {
		relayerCount := viper.GetInt(fmt.Sprintf("%s.relayer_count", chain))
		accounts[chain] = make(chan *models.AccountAggregate, relayerCount)
	}

	return service
}

func (serv *accountService) GetAccountForChain(ctx context.Context, chain string) {

}

func (serv *accountService) name() {

}

func (serv *accountService) fundAccounts() {

}
