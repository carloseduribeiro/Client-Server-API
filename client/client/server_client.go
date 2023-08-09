package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	cotacaoResource        = "http://localhost:8080/cotacao"
	cotacaoResourceTimeout = 300 * time.Millisecond
)

func GetExchange(ctx context.Context) (*ExchangeDto, error) {
	ctx, cancel := context.WithTimeout(ctx, cotacaoResourceTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cotacaoResource, http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro na chamada para a api: %s", err)
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	if response.StatusCode != http.StatusOK {
		var messageResponse *ResponseDto
		if err = decoder.Decode(messageResponse); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(messageResponse.Message)
	}
	var responseData []ExchangeDto
	if err = decoder.Decode(&responseData); err != nil {
		return nil, err
	}
	return &responseData[0], nil
}
