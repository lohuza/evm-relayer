package ports

import (
	"context"

	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/pkg/crud"
)

type AccountService interface {
}

type AccountRepository interface {
	crud.Crud[models.Account]
	GetAvailableAccountsForChain(ctx context.Context, chain string, accountAmount int32) ([]models.Account, error)
}
