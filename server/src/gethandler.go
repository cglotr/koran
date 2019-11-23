package main

import (
	"log"
	"net/http"
)

func getHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.RawPath)
		w.Header().Add("Content-Type", "application/json;charset=utf-8")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h(w, r)
	}
}
