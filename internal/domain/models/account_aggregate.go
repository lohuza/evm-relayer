package models

type AccountAggregate struct {
	Account
	Nonce int64
}

func NewAccountAggregate(account Account, nonce int64) *AccountAggregate {
	return &AccountAggregate{
		Account: account,
		Nonce:   nonce,
	}
}

func (acc *AccountAggregate) IncrementNonce() {
	acc.Nonce++
}
