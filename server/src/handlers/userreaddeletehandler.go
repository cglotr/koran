package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cglotr/koran/server/src/database"
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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		rb := RequestBody{}
		json.NewDecoder(r.Body).Decode(&rb)
		err = u.DeleteUserQuran(userID, rb.SuraID, rb.VerseID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
	}
}
