package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
)

// SuraHandler .
func SuraHandler(q database.QuranGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		suraNumber, err := utils.GetIntURLQuery(r.URL.Query(), "suraNumber")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		verseNumber, err := utils.GetIntURLQuery(r.URL.Query(), "verseNumber")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		verses, err := q.GetVerses(suraNumber, verseNumber, 1)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(payload{
				Message: err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(verses)
	}
}