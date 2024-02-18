package ports

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models/account"
)

type BlockchainService interface {
	GetNonceForAccount(ctx context.Context, account account.Account) (uint64, error)
	GetNoncesForAccounts(ctx context.Context, accounts []account.Account) (account.AccountIDToNonceMap, error)
}
