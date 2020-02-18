package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/constants"
	"github.com/cglotr/koran/server/src/user"
	"github.com/cglotr/koran/server/src/utils"
)

// AuthHandler .
func AuthHandler(c *auth.Client, u user.CRUD) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			IDToken string `json:"id_token"`
		}

		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)

		idToken := rb.IDToken
		t, err := c.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		if !u.Exist(t.UID) {
			log.Printf("user doesn't exist. creating %v\n", t.UID)
			u.Create(t.UID)
		}

		user, err := u.Get(t.UID)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		token := idToken[len(idToken)-constants.AuthTokenLength:]
		u.UpdateToken(user.ID, token)
		user, err = u.Get(t.UID)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
