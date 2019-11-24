package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuraHandler(t *testing.T) {
	testSuraQuery(t, "", http.StatusBadRequest)
	testSuraQuery(t, "suraNumber=1", http.StatusBadRequest)
	testSuraQuery(t, "verseNumber=1", http.StatusBadRequest)
	testSuraQuery(t, "suraNumber=1&verseNumber=1", http.StatusOK)
	testSuraQuery(t, "suraNumber=0&verseNumber=0", http.StatusInternalServerError)
	testSuraQuery(t, "suraNumber=0&verseNumber=1", http.StatusInternalServerError)
	testSuraQuery(t, "suraNumber=1&verseNumber=0", http.StatusInternalServerError)
	testSuraResponse(t)
}

func testSuraResponse(t *testing.T) {
	r := httptest.NewRequest("GET", "http://example.com/sura?suraNumber=1&verseNumber=1", nil)
	q := &stub{}
	w := httptest.NewRecorder()
	suraHandler(q)(w, r)

	var gots []verse
	var wants []verse

	json.NewDecoder(w.Result().Body).Decode(&gots)
	wants, err := q.getVerses(1, 1, 1)
	if err != nil {
		t.FailNow()
	}
	if len(gots) < 1 {
		t.FailNow()
	}
	for i := 0; i < len(wants); i++ {
		got := gots[i]
		want := wants[i]
		if got != want {
			t.Errorf("got %#v; want %#v", got, want)
		}
	}
}

func testSuraQuery(t *testing.T, query string, statusCode int) {
	r := httptest.NewRequest("GET", "http://example.com/sura?"+query, nil)
	q := &stub{}
	w := httptest.NewRecorder()
	suraHandler(q)(w, r)
	assertInt(t, w.Result().StatusCode, statusCode)
}
