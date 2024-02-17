package postgres

import (
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg/crud"
	"github.com/uptrace/bun"
)

type transactionRepository struct {
	crud.Crud[models.Transaction]
}

func NewTransactionRepository(db bun.IDB) ports.TransactionRepository {
	return &transactionRepository{crud.NewCrud[models.Transaction](db)}
}
