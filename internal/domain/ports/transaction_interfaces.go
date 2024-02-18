package ports

import (
	"github.com/lohuza/relayer/internal/domain/models/transaction"
	"github.com/lohuza/relayer/pkg/crud"
)

type TransactionService interface {
}

type TransactionRepository interface {
	crud.Crud[transaction.Transaction]
}
