package account

type ChainToAccountQueue map[string]chan *AccountAggregate

type AccountAggregate struct {
	account Account
	nonce   uint64
}

func NewAccountAggregate(account Account, nonce uint64) AccountAggregate {
	return AccountAggregate{
		account: account,
		nonce:   nonce,
	}
}

func (acc *AccountAggregate) IncrementNonce() {
	acc.nonce++
}

func (acc *AccountAggregate) GetNonce() uint64 {
	return acc.nonce
}
