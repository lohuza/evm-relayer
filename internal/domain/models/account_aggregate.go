package models

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

var (
	ErrInvalidNonces = errors.New("invalid nonces for creating account aggregate")
)

type AccountIDToNonceMap map[xid.ID]int64

type AccountAggregate struct {
	accounts      []Account
	accountNonces AccountIDToNonceMap
	mutex         *sync.Mutex
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
		mutex:         &sync.Mutex{},
	}, nil
}

func (acc *AccountAggregate) IncrementNonce(accountID xid.ID) {
	acc.mutex.Lock()
	acc.accountNonces[accountID]++
	acc.mutex.Unlock()
}
