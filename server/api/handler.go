package api

import (
	"context"
	"errors"
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
	"net/http"
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
	exchangeResult, err := e.getExchange(r.Context(), coins)
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
	err = e.repository.Save(r.Context(), &exchangeResult[0])
	if err != nil {
		writeResponseErr(w, http.StatusInternalServerError, err)
		return
	}
	_ = writeJSON(w, http.StatusOK, exchangeResult)
}
