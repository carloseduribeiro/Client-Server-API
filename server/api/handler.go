package api

import (
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
	"net/http"
)

const (
	coins = "USD-BRL"
)

type ExchangeHandler struct {
	getExchange exchange.Getter
	repository  ExchangeRepository
}

func NewExchangeHandler(getExchange exchange.Getter, repository ExchangeRepository) *ExchangeHandler {
	return &ExchangeHandler{
		getExchange: getExchange,
		repository:  repository,
	}
}

func (e *ExchangeHandler) GetExchange(w http.ResponseWriter, r *http.Request) {
	exchangeResult, err := e.getExchange(r.Context(), coins)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	err = e.repository.Save(r.Context(), exchangeResult)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	_ = WriteJSON(w, http.StatusOK, exchangeResult)
}
