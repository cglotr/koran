package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
	"github.com/gorilla/mux"
)

// UserReadHandler .
func UserReadHandler(u database.UserQuranReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type ResponseBody struct {
			Read bool `json:"read"`
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
		suraID, err := utils.GetIntURLQuery(r.URL.Query(), "sura_id")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
		verseID, err := utils.GetIntURLQuery(r.URL.Query(), "verse_id")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		ok, err := u.ReadUserQuran(userID, suraID, verseID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}
		responseBody := ResponseBody{
			Read: ok,
		}
		json.NewEncoder(w).Encode(responseBody)
	}
}
