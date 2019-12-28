package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/koran/server/src/jsons"
	"github.com/cglotr/koran/server/src/utils"
)

func TestRootHandler(t *testing.T) {
	m := jsons.Message{}
	r := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	RootHandler()(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.FailNow()
	}
	json.NewDecoder(w.Result().Body).Decode(&m)
	utils.AssertString(t, m.Message, "👋👋 Welcome to Koran!")
}
