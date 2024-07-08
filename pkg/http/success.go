package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string            `json:"message"`
	Data    map[string]string `json:"data,omitempty"`
	Success bool              `json:"success"`
}

func WriteSuccessJson(message string, data map[string]string, success bool, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Message: message,
		Data:    data,
		Success: success,
	})
}
