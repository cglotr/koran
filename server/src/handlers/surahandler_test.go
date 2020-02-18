package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/koran/server/src/database"
	"github.com/cglotr/koran/server/src/handlers"
	"github.com/cglotr/koran/server/src/structs"
	"github.com/cglotr/koran/server/src/utils"
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
	q := &database.Stub{}
	w := httptest.NewRecorder()
	SuraHandler(q)(w, r)

	var gots []structs.Verse
	var wants []structs.Verse

	json.NewDecoder(w.Result().Body).Decode(&gots)
	wants, err := q.GetVerses(1, 1, 1)
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
	q := &database.Stub{}
	w := httptest.NewRecorder()
	handlers.SuraHandler(q)(w, r)
	utils.AssertInt(t, w.Result().StatusCode, statusCode)
}
