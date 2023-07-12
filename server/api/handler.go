package api

import (
	"context"
	"encoding/json"
	"github.com/carloseduribeiro/Client-Server-API/server/client"
	"net/http"
	"time"
)

const (
	coins                      = "USD-BRL"
	exchangeAPITimeout         = 200 * time.Millisecond
	databasePersistenceTimeout = 10 * time.Millisecond
)

type ExchangeClient interface {
	GetExchange(ctx context.Context, coins string) (*client.ExchangeDto, error)
}

type ExchangeRepository interface {
	Save(ctx context.Context, dto *client.ExchangeDto) error
}

type ExchangeHandler struct {
	client     ExchangeClient
	repository ExchangeRepository
}

type ResponseDto struct {
	Message string `json:"message"`
}

func NewExchangeHandler(client ExchangeClient, repository ExchangeRepository) *ExchangeHandler {
	return &ExchangeHandler{
		client:     client,
		repository: repository,
	}
}

func (e *ExchangeHandler) GetExchange(w http.ResponseWriter, r *http.Request) {
	requestCtx, cancelReq := context.WithTimeout(r.Context(), exchangeAPITimeout)
	defer cancelReq()
	exchange, err := e.client.GetExchange(requestCtx, coins)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	persistCtx, cancelPersist := context.WithTimeout(r.Context(), databasePersistenceTimeout)
	defer cancelPersist()
	err = e.repository.Save(persistCtx, exchange)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	_ = WriteJSON(w, http.StatusOK, exchange)
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	statusCode := http.StatusGatewayTimeout
	if err != context.DeadlineExceeded {
		statusCode = http.StatusInternalServerError
	}
	_ = WriteJSON(w, statusCode, ResponseDto{Message: err.Error()})
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
