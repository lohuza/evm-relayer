package ports

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models"
)

type BlockchainService interface {
	GetNonceForAccount(ctx context.Context, account models.Account) (uint64, error)
	GetNoncesForAccounts(ctx context.Context, accounts []models.Account) (models.AccountIDToNonceMap, error)
}
