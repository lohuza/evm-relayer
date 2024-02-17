package ports

import (
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/pkg/crud"
)

type TransactionService interface {
}

type TransactionRepository interface {
	crud.Crud[models.Transaction]
}
