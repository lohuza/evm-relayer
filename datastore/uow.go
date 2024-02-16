package datastore

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type UnitOfWork interface {
	RunInTx(context.Context, UnitOfWorkBlock) error
	DB() UnitOfWorkStore
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

func (s *unitOfWork) DB() UnitOfWorkStore {
	return newUowStore(s.conn)
}
