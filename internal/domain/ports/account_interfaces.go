package ports

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models/account"
	"github.com/lohuza/relayer/pkg/crud"
)

type AccountService interface {
}

type AccountRepository interface {
	crud.Crud[account.Account]
	GetAvailableAccountsForChain(ctx context.Context, chain string, accountAmount int32) ([]account.Account, error)
}
