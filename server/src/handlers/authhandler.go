package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/constants"
)

// AuthHandler .
func AuthHandler(c *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		idToken := values["id_token"][0]
		token, err := c.VerifyIDToken(context.Background(), idToken)
		log.Println(fmt.Sprintf("%v", token))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(payload{
			Token: idToken[:constants.AuthTokenLength],
			UID:   token.UID,
		})
	}
}
