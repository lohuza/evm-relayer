package repository

import (
	"github.com/lohuza/relayer/datastore/model"
	"github.com/uptrace/bun"
)

type AccountRepository interface {
	RepositoryBase[model.Account]
}

type accountRepository struct {
	RepositoryBase[model.Account]
}

func NewAccountRepository(db bun.IDB) AccountRepository {
	return &accountRepository{newBaseRepository[model.Account](db)}
}
