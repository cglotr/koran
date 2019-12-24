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
func AuthHandler(c *auth.Client, u database.UserCRUD) http.HandlerFunc {
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
			log.Printf("user doesn't exist. creating %v\n", t.UID)
			u.CreateUser(t.UID)
		}
		user, err := u.GetUser(t.UID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		token := idToken[:constants.AuthTokenLength]
		u.UpdateUserToken(user.ID, token)
		user, err = u.GetUser(t.UID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
