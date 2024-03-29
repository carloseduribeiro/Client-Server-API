package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	exchangeAPIResource = "https://economia.awesomeapi.com.br/json/last/"
	exchangeAPITimeout  = 200 * time.Millisecond
)

type GetExchangeFunc func(ctx context.Context, coins string) ([]Exchange, error)

func GetExchange(ctx context.Context, coins string) ([]Exchange, error) {
	reqCtx, cancel := context.WithTimeout(ctx, exchangeAPITimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, exchangeAPIResource+coins, http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro na chamada para a api: %s\n", err)
		return nil, err
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	if response.StatusCode == http.StatusNotFound {
		responseBody := ErrorDto{}
		if err = decoder.Decode(&responseBody); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(responseBody.Message)
	}
	baseResponse := make(SuccessDto)
	if err = decoder.Decode(&baseResponse); err != nil {
		return nil, err
	}
	if exchange, ok := baseResponse[strings.ReplaceAll(coins, "-", "")]; ok {
		return []Exchange{exchange}, nil
	}
	return nil, fmt.Errorf("erro no retorno da api de câmbio")
}
