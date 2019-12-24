package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/constants"
	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
)

// AuthHandler .
func AuthHandler(c *auth.Client, u database.UserCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToken, err := utils.GetStringURLQuery(r.URL.Query(), "id_token")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		t, err := c.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !u.IsUserExist(t.UID) {
			log.Printf("user doesn't exist. creating %v", t.UID)
			u.CreateUser(t.UID)
		}
		token := idToken[:constants.AuthTokenLength]
		json.NewEncoder(w).Encode(payload{
			Token: token,
			UID:   t.UID,
		})
	}
}
