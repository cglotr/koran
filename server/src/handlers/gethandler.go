package handlers

import (
	"net/http"
)

// GetHandler .
func GetHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		BaseHandler(h)(w, r)
	}
}
