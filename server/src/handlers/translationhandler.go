package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/utils"
)

// TranslationHandler .
func TranslationHandler(t database.TranslationsGetter) http.HandlerFunc {
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

		translations, err := t.GetTranslations("english-yusuf-al", suraNumber, verseNumber, 1)
		if err != nil {
			utils.WriteMessageError(w, http.StatusInternalServerError, err)
			return
		}

		json.NewEncoder(w).Encode(translations)
	}
}
