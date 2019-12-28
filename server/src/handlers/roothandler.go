package handlers

import (
	"net/http"

	"github.com/cglotr/koran/server/src/utils"
)

// RootHandler .
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteMessage(w, http.StatusOK, "👋👋 Welcome to Koran!")
	}
}
