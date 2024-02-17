package ports

import (
	"github.com/lohuza/relayer/internal/infrastructure/httpclient/dto"
)

type GasPriceService interface {
}

type GasPriceClient interface {
	FetchGasPrices(networkID int32) (*dto.GasPriceResponse, error)
}
