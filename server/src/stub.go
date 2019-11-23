package main

import (
	"errors"
	"fmt"
)

type stub struct{}

func (s *stub) getVerses(suraNumber, startVerse, numberOfVerses int) ([]verse, error) {
	if suraNumber < 1 {
		return nil, errors.New("invalid suraNumber")
	}
	if startVerse < 1 {
		return nil, errors.New("invalid startVerse")
	}
	if numberOfVerses < 1 {
		return nil, errors.New("invalid numberOfVerses")
	}
	verses := []verse{}
	for i := 0; i < numberOfVerses; i++ {
		verses = append(verses, verse{
			Ayah:        fmt.Sprintf("sura %v, verse %v", suraNumber, startVerse+i),
			SuraNumber:  suraNumber,
			VerseNumber: startVerse + i,
		})
	}
	return verses, nil
}
