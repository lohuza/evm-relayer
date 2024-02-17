package ports

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/pkg/crud"
)

type AccountService interface {
	CreateAccounts(ctx context.Context, chain string, accountCount int32) ([]*models.AccountAggregate, error)
}

type AccountRepository interface {
	crud.Crud[models.Account]
}
