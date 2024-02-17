package postgres

import (
	"context"
	"database/sql"

	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/uptrace/bun"
)

type UnitOfWorkBlock func(UnitOfWorkStore) error

type UnitOfWorkStore interface {
	Account() ports.AccountRepository
	Transaction() ports.TransactionRepository
}

type unitOfWorkStore struct {
	accountRepo     ports.AccountRepository
	transactionRepo ports.TransactionRepository
}

func (u *unitOfWorkStore) Account() ports.AccountRepository {
	return u.accountRepo
}

func (u *unitOfWorkStore) Transaction() ports.TransactionRepository {
	return u.transactionRepo
}

func newUowStore(db bun.IDB) UnitOfWorkStore {
	return &unitOfWorkStore{
		accountRepo:     NewAccountRepository(db),
		transactionRepo: NewTransactionRepository(db),
	}
}

type UnitOfWork interface {
	RunInTx(context.Context, UnitOfWorkBlock) error
	Repo() UnitOfWorkStore
}

type unitOfWork struct {
	conn *bun.DB
}

func NewStore(db *bun.DB) UnitOfWork {
	return &unitOfWork{
		conn: db,
	}
}

func (s *unitOfWork) RunInTx(ctx context.Context, fn UnitOfWorkBlock) error {
	return s.conn.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		uow := newUowStore(tx)
		return fn(uow)
	})
}

func (s *unitOfWork) Repo() UnitOfWorkStore {
	return newUowStore(s.conn)
}
