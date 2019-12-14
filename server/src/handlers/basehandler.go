package handlers

import (
	"log"
	"net/http"
)

// BaseHandler .
func BaseHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + r.URL.RequestURI())
		w.Header().Add("Content-Type", "application/json;charset=utf-8")
		h(w, r)
	}
}
