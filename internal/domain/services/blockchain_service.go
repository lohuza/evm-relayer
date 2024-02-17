package services

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/rs/zerolog/log"
)

var (
	ErrProviderDoesNotExist = errors.New("provider for provided chain doesn't exist")
	ErrCastingPublicKet     = errors.New("error casting public key to ECDSA")
)

type blockchainService struct {
	providers map[string]*ethclient.Client
	store     postgres.UnitOfWork
}

func NewBlockchainService() ports.BlockchainService {
	return &blockchainService{}
}

func (service *blockchainService) GetNoncesForAccounts(ctx context.Context, accounts []models.Account) (map[int32]int64, error) {

}

func (service *blockchainService) GetNonceForAccount(ctx context.Context, account models.Account) (int64, error) {
	provider, exists := service.providers[account.Chain]
	if !exists {
		return 0, ErrProviderDoesNotExist
	}

	provider.PendingNonceAt(ctx)
}

func (service *blockchainService) GetAddressFromPrivateKey(privateKeyHex string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {

	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Err(err).Msg("Error casting public key to ECDSA")
		return common.Address{}, ErrCastingPublicKet
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
