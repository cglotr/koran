package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuraHandler(t *testing.T) {
	testSuraQuery(t, "", http.StatusBadRequest)
	testSuraQuery(t, "suraNumber=1", http.StatusBadRequest)
	testSuraQuery(t, "verseNumber=1", http.StatusBadRequest)
	testSuraQuery(t, "suraNumber=1&verseNumber=1", http.StatusOK)
}

func testSuraQuery(t *testing.T, query string, statusCode int) {
	r := httptest.NewRequest("GET", "http://example.com/sura?"+query, nil)
	q := &stub{}
	w := httptest.NewRecorder()
	suraHandler(q)(w, r)
	assertInt(t, w.Result().StatusCode, statusCode)
}
