package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

func WriteErrorJson(message string, error string, success bool, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
		Error:   error,
		Success: success,
	})
}
