package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielbsx/wyd-go/app"
)

func sendError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(app.ErrorResponse{
		Message: msg,
		Code:    status,
	})
}

func sendSuccess(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
