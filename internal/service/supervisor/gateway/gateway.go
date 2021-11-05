package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/foxfurry/go_aggregator/internal/domain/dto"
	"net/http"
)

type IGateway interface {
	SendOrder(order *dto.Order, host string) (*http.Response, error)
}

type aggregatorGateway struct {

}

func NewAggregatorGateway() IGateway{
	return &aggregatorGateway{}
}

func (g *aggregatorGateway) SendOrder(order *dto.Order, host string) (*http.Response, error) {
	jsonBody, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	contentType := "application/json"
	return http.Post(host + "/order", contentType, bytes.NewReader(jsonBody))
}
