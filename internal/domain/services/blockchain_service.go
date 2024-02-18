package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lohuza/relayer/internal/adapters/repository/postgres"
	"github.com/lohuza/relayer/internal/domain/models/account"
	"github.com/lohuza/relayer/internal/domain/models/transaction"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg"
	"github.com/rs/zerolog/log"
)

var (
	ErrProviderDoesNotExist = errors.New("provider for provided chain doesn't exist")
	ErrCastingPublicKet     = errors.New("error casting public key to ECDSA")
)

type ChainToRpc map[string]*ethclient.Client
type ChainToChainID map[string]*big.Int

type blockchainService struct {
	// chain to client
	providers ChainToRpc
	store     postgres.UnitOfWork
	mutex     *sync.RWMutex
}

func NewBlockchainService() ports.BlockchainService {
	return &blockchainService{
		providers: nil,
		store:     nil,
		mutex:     &sync.RWMutex{},
	}
}

func (service *blockchainService) GetNoncesForAccounts(ctx context.Context, accounts []account.Account) (account.AccountIDToNonceMap, error) {
	var accountIDToNonceMap account.AccountIDToNonceMap
	for _, account := range accounts {
		nonce, err := pkg.ExecuteWithRetry(func() (uint64, error) {
			return service.GetNonceForAccount(ctx, account)
		})
		if err != nil {
			return nil, err
		}
		accountIDToNonceMap[account.ID] = nonce
	}

	return accountIDToNonceMap, nil
}

func (service *blockchainService) GetNonceForAccount(ctx context.Context, account account.Account) (uint64, error) {
	provider, exists := service.providers[account.Chain]
	if !exists {
		return 0, ErrProviderDoesNotExist
	}

	address, err := service.GetAddressFromPrivateKey(account.PrivateKey)
	if err != nil {
		return 0, err
	}
	nonce, err := provider.PendingNonceAt(ctx, address)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get nonce for account id %v", account.ID)
		return 0, err
	}
	return nonce, nil
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

func (service *blockchainService) SendTransaction(ctx context.Context, transaction transaction.Transaction) {
	client := service.getRpcClient(transaction.Chain)
	//chainID, err := client.ChainID()
	//types.NewTx()
	chainID := big.NewInt(123)
	types.NewEIP155Signer(chainID)
}

func (service *blockchainService) getRpcClient(chain string) *ethclient.Client {
	service.mutex.RLock()
	defer service.mutex.RUnlock()

	return service.providers[chain]
}

func (service *blockchainService) addRpcClient(chain string, client *ethclient.Client) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.providers[chain] = client
}

func (s *blockchainService) name() {

}
