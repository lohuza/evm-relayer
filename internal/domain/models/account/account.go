package account

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/xid"
	"github.com/uptrace/bun"
)

type AccountIDToNonceMap map[xid.ID]uint64

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID         xid.ID `bun:"id,pk"`
	PrivateKey string `bun:"private_key"`
	Chain      string `bun:"chain"`
	IsInUse    bool   `bun:"is_in_use"`
}

func NewAccount(chain string, isInUse bool) (*Account, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	return &Account{
		ID:         xid.New(),
		PrivateKey: privateKeyHex,
		Chain:      chain,
		IsInUse:    isInUse,
	}, nil
}

func (acc *Account) MarkAsBeingUsed() {
	acc.IsInUse = true
}

func (acc *Account) MarkAsBeingAvailable() {
	acc.IsInUse = false
}
