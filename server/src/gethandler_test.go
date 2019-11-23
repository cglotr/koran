package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	testGetRequest(t)
	testNonGetRequest(t)
}

func testGetRequest(t *testing.T) {
	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	getHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})(w, r)

	assertString(t, w.Header().Get("Content-Type"), "application/json;charset=utf-8")
	assertString(t, w.Body.String(), "test")
}

func testNonGetRequest(t *testing.T) {
	r := httptest.NewRequest("NOTGET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	getHandler(func(w http.ResponseWriter, r *http.Request) {
		t.Fail()
	})(w, r)

	assertInt(t, w.Result().StatusCode, http.StatusBadRequest)
}
