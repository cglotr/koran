package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
	"github.com/gorilla/mux"
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
			utils.WriteMessageError(w, http.StatusBadRequest, err)
			return
		}

		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)
		token := rb.Token

		user, err := u.GetUserByID(id)
		if err != nil {
			utils.WriteMessageError(w, http.StatusNoContent, err)
			return
		}

		if user.Token != "" && user.Token != token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		u.UpdateUserToken(user.ID, "")
		utils.WriteMessage(w, http.StatusOK, "👌")
	}
}
