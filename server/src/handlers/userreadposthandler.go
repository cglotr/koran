package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
	"github.com/gorilla/mux"
)

// UserReadPostHandler .
func UserReadPostHandler(u database.UserQuranCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			SuraID  int `json:"sura_id"`
			VerseID int `json:"verse_id"`
		}

		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.WriteMessageError(w, http.StatusBadRequest, err)
			return
		}

		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)
		u.CreateUserQuran(userID, rb.SuraID, rb.VerseID)
		utils.WriteMessage(w, http.StatusOK, "👌")
	}
}
