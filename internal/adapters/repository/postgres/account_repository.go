package postgres

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models/account"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg/crud"
	"github.com/uptrace/bun"
)

type accountRepository struct {
	crud.Crud[account.Account]
}

func NewAccountRepository(db bun.IDB) ports.AccountRepository {
	return &accountRepository{crud.NewCrud[account.Account](db)}
}

func (repo *accountRepository) GetAvailableAccountsForChain(ctx context.Context, chain string, accountAmount int32) ([]account.Account, error) {
	return repo.FindAll(ctx, func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("chain = ? AND is_in_use", chain, false).Limit(int(accountAmount))
	})
}
