package utils

import (
	"encoding/json"
	"net/http"

	"github.com/cglotr/koran/server/src/jsons"
)

// WriteMessage .
func WriteMessage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(jsons.Message{
		Message: message,
	})
}
