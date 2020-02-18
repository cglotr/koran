package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/structs"
)

// WriteMessageError .
func WriteMessageError(w http.ResponseWriter, statusCode int, err error) {
	log.Println(err.Error())
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(structs.Message{
		Message: err.Error(),
	})
}
