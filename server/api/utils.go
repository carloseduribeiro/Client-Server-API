package api

import (
	"encoding/json"
	"net/http"
)

func writeResponseErr(w http.ResponseWriter, code int, err error) {
	_ = writeJSON(w, code, ResponseDto{Message: err.Error()})
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
