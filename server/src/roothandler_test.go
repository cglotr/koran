package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	p := payload{}
	r := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	rootHandler()(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.FailNow()
	}
	json.NewDecoder(w.Result().Body).Decode(&p)
	assertString(t, p.Message, "👋👋 Welcome to Koran!")
}
