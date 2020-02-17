package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cglotr/koran/server/src/quran"
	"github.com/cglotr/koran/server/src/utils"
)

// SuraHandler .
func SuraHandler(q quran.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		suraNumber, err := utils.GetIntURLQuery(r.URL.Query(), "suraNumber")
		if err != nil {
			utils.WriteMessageError(w, http.StatusBadRequest, err)
			return
		}

		verseNumber, err := utils.GetIntURLQuery(r.URL.Query(), "verseNumber")
		if err != nil {
			utils.WriteMessageError(w, http.StatusBadRequest, err)
			return
		}

		verses, err := q.GetVerses(suraNumber, verseNumber, 1)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		json.NewEncoder(w).Encode(verses)
	}
}
