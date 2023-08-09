package api

import (
	"context"
	"errors"
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
	"log"
	"net/http"
	"time"
)

const (
	coins = "USD-BRL"
)

type ExchangeHandler struct {
	getExchange exchange.GetExchangeFunc
	repository  ExchangeRepository
}

func NewExchangeHandler(getExchangeFunc exchange.GetExchangeFunc, repository ExchangeRepository) *ExchangeHandler {
	return &ExchangeHandler{
		getExchange: getExchangeFunc,
		repository:  repository,
	}
}

func (e *ExchangeHandler) GetExchange(w http.ResponseWriter, r *http.Request) {
	log.Println("iniciarndo request:", r.Method, r.RequestURI)
	defer log.Println("request finalizada")
	ctx := r.Context()
	select {
	case <-time.After(200 * time.Millisecond):
		exchangeResult, err := e.getExchange(ctx, coins)
		if err != nil {
			if exchangeResult != nil && len(exchangeResult) == 0 {
				writeResponseErr(w, http.StatusNotFound, err)
				return
			}
			if errors.Is(err, context.DeadlineExceeded) {
				writeResponseErr(w, http.StatusRequestTimeout, err)
				return
			}
			writeResponseErr(w, http.StatusInternalServerError, err)
			return
		}
		err = e.repository.Save(ctx, &exchangeResult[0])
		if err != nil {
			writeResponseErr(w, http.StatusInternalServerError, err)
			return
		}
		_ = writeJSON(w, http.StatusOK, exchangeResult)
	case <-ctx.Done():
		log.Println("request finalizada pelo usuário")
		return
	}
}
