package services

import "github.com/lohuza/relayer/internal/domain/ports"

type gasPriceService struct {
}

func NewGasPriceService() ports.GasPriceService {
	return &gasPriceService{}
}

func (service *gasPriceService) SetGasPrice(chain string, price float32, unit string) {

}
