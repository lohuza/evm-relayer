package models

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

var (
	ErrInvalidNonces = errors.New("invalid nonces for creating account aggregate")
)

type AccountIDToNonceMap map[xid.ID]uint64

type AccountAggregate struct {
	accounts      []Account
	accountNonces AccountIDToNonceMap
	lock          *sync.RWMutex
}

func NewAccountAggregate(accounts []Account, accountNonces AccountIDToNonceMap) (*AccountAggregate, error) {
	if len(accounts) != len(accountNonces) {
		return nil, ErrInvalidNonces
	}

	for _, account := range accounts {
		if _, exists := accountNonces[account.ID]; !exists {
			return nil, ErrInvalidNonces
		}
	}

	return &AccountAggregate{
		accounts:      accounts,
		accountNonces: accountNonces,
		lock:          &sync.RWMutex{},
	}, nil
}

func (acc *AccountAggregate) IncrementNonce(accountID xid.ID) {
	acc.lock.Lock()
	defer acc.lock.Unlock()

	acc.accountNonces[accountID]++
}

func (acc *AccountAggregate) GetNonce(accountID xid.ID) uint64 {
	acc.lock.RLock()
	defer acc.lock.RUnlock()

	return acc.accountNonces[accountID]
}
