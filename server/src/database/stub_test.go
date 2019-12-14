package database

import (
	"fmt"
	"testing"
)

func TestGetVerses(t *testing.T) {
	s := Stub{}
	valids := [][3]int{
		[3]int{1, 1, 1},
		[3]int{2, 1, 1},
		[3]int{2, 2, 1},
		[3]int{2, 2, 10},
	}
	for _, test := range valids {
		suraNumber := test[0]
		startVerse := test[1]
		numberOfVerses := test[2]
		verses, err := s.GetVerses(suraNumber, startVerse, numberOfVerses)
		if err != nil {
			t.Fail()
		}
		if len(verses) != numberOfVerses {
			t.Fail()
		}
		for _, verse := range verses {
			got := verse.SuraNumber
			want := suraNumber
			if got != want {
				t.Errorf("got %#v; want %#v", got, want)
			}
		}
		for i, verse := range verses {
			got := verse.VerseNumber
			want := startVerse + i
			if got != want {
				t.Errorf("got %#v; want %#v", got, want)
			}
		}
		for i, verse := range verses {
			got := verse.Ayah
			want := fmt.Sprintf("sura %v, verse %v", suraNumber, startVerse+i)
			if got != want {
				t.Errorf("got %#v; want %#v", got, want)
			}
		}
	}
	invalids := [][3]int{
		[3]int{-1, -1, -1},
		[3]int{0, 0, 0},
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
		[3]int{1, 1, 0},
		[3]int{0, 1, 1},
	}
	for _, test := range invalids {
		suraNumber := test[0]
		startVerse := test[1]
		numberOfVerses := test[2]
		verses, err := s.GetVerses(suraNumber, startVerse, numberOfVerses)
		if verses != nil {
			t.Fail()
		}
		if err == nil {
			t.Errorf("want %#v, %#v, %#v to fail", suraNumber, startVerse, numberOfVerses)
		}
	}
}
