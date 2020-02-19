package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/cglotr/koran/server/src/user"
	"github.com/cglotr/koran/server/src/utils"
	"github.com/gorilla/mux"
)

// AuthInvalidateHandler .
func AuthInvalidateHandler(c *auth.Client, u user.CRUD) http.HandlerFunc {
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

		user, err := u.GetByID(id)
		if err != nil {
			utils.WriteMessageError(w, http.StatusNoContent, err)
			return
		}

		if user.Token != "" && user.Token != token {
			utils.WriteMessage(w, http.StatusUnauthorized, "🚫")
			return
		}
		u.UpdateToken(user.ID, "")
		utils.WriteMessage(w, http.StatusOK, "👌")
	}
}
