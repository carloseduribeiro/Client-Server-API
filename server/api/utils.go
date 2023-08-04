package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	statusCode := http.StatusGatewayTimeout
	if !errors.Is(err, context.DeadlineExceeded) {
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
