package repository

import (
	"github.com/lohuza/relayer/datastore/model"
	"github.com/uptrace/bun"
)

type TransactionRepository interface {
	RepositoryBase[model.Transaction]
}

type transactionRepository struct {
	RepositoryBase[model.Transaction]
}

func NewTransactionRepository(db bun.IDB) TransactionRepository {
	return &transactionRepository{newBaseRepository[model.Transaction](db)}
}
