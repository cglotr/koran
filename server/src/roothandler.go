package main

import (
	"encoding/json"
	"net/http"
)

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(payload{
			Message: "👋👋 Welcome to Koran!",
		})
	}
}
