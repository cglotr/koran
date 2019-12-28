package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/jsons"
)

// WriteMessageError .
func WriteMessageError(w http.ResponseWriter, statusCode int, err error) {
	log.Println(err.Error())
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(jsons.Message{
		Message: err.Error(),
	})
}
