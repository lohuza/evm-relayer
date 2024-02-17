package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lohuza/relayer/internal/domain/ports"
	"github.com/lohuza/relayer/internal/infrastructure/httpclient/dto"
)

const gasPriceApiUrl = "https://api.biconomy.io/api/v1/gas-price?networkId="

type gasPriceClient struct {
	client http.Client
}

func NewGasPriceClient() ports.GasPriceClient {
	return &gasPriceClient{
		client: http.Client{},
	}
}

func (api *gasPriceClient) FetchGasPrices(networkID int32) (*dto.GasPriceResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s%v", gasPriceApiUrl, networkID))
	if err != nil {
		return nil, fmt.Errorf("error fetching gas price: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var response dto.GasPriceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return &response, nil
}
