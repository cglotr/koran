package handlers

import (
	"encoding/json"
	"net/http"
)

// RootHandler .
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(payload{
			Message: "👋👋 Welcome to Koran!",
		})
	}
}
