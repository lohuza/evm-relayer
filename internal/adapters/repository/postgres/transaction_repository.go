package postgres

import (
	"github.com/lohuza/relayer/internal/domain/models/transaction"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg/crud"
	"github.com/uptrace/bun"
)

type transactionRepository struct {
	crud.Crud[transaction.Transaction]
}

func NewTransactionRepository(db bun.IDB) ports.TransactionRepository {
	return &transactionRepository{crud.NewCrud[transaction.Transaction](db)}
}
