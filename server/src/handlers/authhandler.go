package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/constants"
	"github.com/cglotr/koran/server/src/database"
)

// AuthHandler .
func AuthHandler(c *auth.Client, u database.UserCRUD) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			IDToken string `json:"id_token"`
		}
		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)
		idToken := rb.IDToken
		t, err := c.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
		if !u.IsUserExist(t.UID) {
			log.Printf("user doesn't exist. creating %v\n", t.UID)
			u.CreateUser(t.UID)
		}
		user, err := u.GetUser(t.UID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
		token := idToken[len(idToken)-constants.AuthTokenLength:]
		u.UpdateUserToken(user.ID, token)
		user, err = u.GetUser(t.UID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
