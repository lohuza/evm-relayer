package services

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/ports"
)

type blockchainService struct {
	providers map[string]*ethclient.Client
	store     postgres.UnitOfWork
}

func NewBlockchainService() ports.BlockchainService {
	return &blockchainService{}
}

func (service *blockchainService) GetNonceForAddress(ctx context.Context, chain string, address string) {

}
