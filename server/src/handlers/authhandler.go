package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
)

// AuthHandler .
func AuthHandler(c *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		idToken := values["id_token"][0]
		token, err := c.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
		}
		json.NewEncoder(w).Encode(payload{Message: fmt.Sprintf("%v", token)})
	}
}
