package dto

type NewTransactionRequest struct {
	To       string `json:"to"`
	Chain    string `json:"chain"`
	Data     string `json:"data"`
	GasLimit string `json:"gas_limit"`
}
