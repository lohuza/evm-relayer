package datastore

import (
	"github.com/lohuza/relayer/datastore/repository"
	"github.com/uptrace/bun"
)

type UnitOfWorkBlock func(UnitOfWorkStore) error

type UnitOfWorkStore interface {
	AccountRepo() repository.AccountRepository
	TransactionRepo() repository.TransactionRepository
}

type unitOfWorkStore struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func (u *unitOfWorkStore) AccountRepo() repository.AccountRepository {
	return u.accountRepo
}

func (u *unitOfWorkStore) TransactionRepo() repository.TransactionRepository {
	return u.transactionRepo
}

func newUowStore(db bun.IDB) UnitOfWorkStore {
	return &unitOfWorkStore{
		accountRepo:     repository.NewAccountRepository(db),
		transactionRepo: repository.NewTransactionRepository(db),
	}
}
