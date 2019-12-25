package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/database"
)

// AuthInvalidateHandler .
func AuthInvalidateHandler(c *auth.Client, u database.UserCRUD) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			Token string `json:"token"`
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)
		token := rb.Token

		user, err := u.GetUserByID(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		if user.Token != token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		u.UpdateUserToken(user.ID, "")
	}
}
