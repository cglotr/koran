package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func suraHandler(q quranGetter) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json;charset=utf-8")

		suraNumber, err := getIntURLQuery(r.URL.Query(), "suraNumber")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		verseNumber, err := getIntURLQuery(r.URL.Query(), "verseNumber")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		verses, err := q.getVerses(suraNumber, verseNumber, 1)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(verses)
		w.WriteHeader(http.StatusOK)
	}
}
