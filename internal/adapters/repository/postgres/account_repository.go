package postgres

import (
	"github.com/lohuza/relayer/internal/domain/models"
	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/pkg/crud"
	"github.com/uptrace/bun"
)

type accountRepository struct {
	crud.Crud[models.Account]
}

func NewAccountRepository(db bun.IDB) ports.AccountRepository {
	return &accountRepository{crud.NewCrud[models.Account](db)}
}
