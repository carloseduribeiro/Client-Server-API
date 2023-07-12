package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiResource = "%s/json/last/%s"
	errMessage  = "moeda nao encontrada %s"
)

type ExchangeDto struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ExchangeClient struct {
	baseURL string
	client  *http.Client
	timeout time.Duration
}

func NewExchangeClient(baseURL string, client *http.Client, timeout time.Duration) *ExchangeClient {
	return &ExchangeClient{
		baseURL: baseURL,
		client:  client,
		timeout: timeout,
	}
}

func (e *ExchangeClient) GetExchange(ctx context.Context, coins string) (*ExchangeDto, error) {
	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()
	url := fmt.Sprintf(apiResource, e.baseURL, coins)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	defer req.Body.Close()
	response, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf(errMessage, coins)
	}
	decoder := json.NewDecoder(response.Body)
	var result ExchangeDto
	if err = decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
