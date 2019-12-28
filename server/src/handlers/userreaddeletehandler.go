package handlers

import (
	"net/http"
	"strconv"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
	"github.com/gorilla/mux"
)

// UserReadDeleteHandler .
func UserReadDeleteHandler(u database.UserQuranDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			SuraID  int `json:"sura_id"`
			VerseID int `json:"verse_id"`
		}

		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		suraID, err := utils.GetIntURLQuery(r.URL.Query(), "sura_id")
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		verseID, err := utils.GetIntURLQuery(r.URL.Query(), "verse_id")
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		err = u.DeleteUserQuran(userID, suraID, verseID)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteMessage(w, http.StatusOK, "👌")
	}
}
