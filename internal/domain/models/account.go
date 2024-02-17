package models

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID         int32  `bun:"id,pk,autoincrement"`
	Address    string `bun:"address"`
	PrivateKey string `bun:"private_key"`
}

func NewAccount() (*Account, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	return &Account{
		Address:    address,
		PrivateKey: privateKeyHex,
	}, nil
}
