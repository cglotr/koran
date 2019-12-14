package handlers

import (
	"net/http"
)

// GetHandler .
func GetHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		BaseHandler(h)(w, r)
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
