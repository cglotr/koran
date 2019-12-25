package handlers

import (
	"net/http"
)

// PostHandler .
func PostHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		BaseHandler(h)(w, r)
	}
}
