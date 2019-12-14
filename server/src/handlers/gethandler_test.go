package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/koran/server/src/utils"
)

func TestGetHandler(t *testing.T) {
	testGetRequest(t)
	testNonGetRequest(t)
}

func testGetRequest(t *testing.T) {
	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	GetHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})(w, r)

	utils.AssertString(t, w.Header().Get("Content-Type"), "application/json;charset=utf-8")
	utils.AssertString(t, w.Body.String(), "test")
}

func testNonGetRequest(t *testing.T) {
	r := httptest.NewRequest("NOTGET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	GetHandler(func(w http.ResponseWriter, r *http.Request) {
		t.Fail()
	})(w, r)

	utils.AssertInt(t, w.Result().StatusCode, http.StatusBadRequest)
}
