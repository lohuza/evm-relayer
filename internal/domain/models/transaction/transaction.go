package transaction

import (
	"github.com/lohuza/relayer/pkg"
	"github.com/uptrace/bun"
)

type TransactionStatus string

const (
	InQueue    TransactionStatus = "in_queue"
	Processing TransactionStatus = "processing"
	Completed  TransactionStatus = "completed"
	Failed     TransactionStatus = "failed"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID              int64             `bun:"id,pk,autoincrement"`
	Chain           string            `bun:"chain"`
	To              string            `bun:"to"`
	Data            string            `bun:"data"`
	GasLimit        string            `bun:"gas_limit"`
	Status          TransactionStatus `bun:"transaction_status"`
	TransactionHash string            `bun:"transaction_hash"`
	CreateTimestamp int64             `bun:"create_timestamp"`
}

func NewTransaction(chain string, to string, data string, gasLimit string) *Transaction {
	return &Transaction{
		Chain:           chain,
		To:              to,
		Data:            data,
		GasLimit:        gasLimit,
		Status:          InQueue,
		CreateTimestamp: pkg.GetUtcTimeSecond(),
	}
}
