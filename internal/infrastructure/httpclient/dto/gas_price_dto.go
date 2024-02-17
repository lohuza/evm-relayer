package dto

type GasPriceResponse struct {
	Code     int32    `json:"code"`
	GasPrice gasPrice `json:"gasPrice"`
}

type gasPrice struct {
	Value int32  `json:"value"`
	Unit  string `json:"unit"`
}
