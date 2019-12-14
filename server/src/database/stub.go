package database

import (
	"errors"
	"fmt"
)

// Stub .
type Stub struct{}

// GetVerses .
func (s *Stub) GetVerses(suraNumber, startVerse, numberOfVerses int) ([]Verse, error) {
	if suraNumber < 1 {
		return nil, errors.New("invalid suraNumber")
	}
	if startVerse < 1 {
		return nil, errors.New("invalid startVerse")
	}
	if numberOfVerses < 1 {
		return nil, errors.New("invalid numberOfVerses")
	}
	verses := []Verse{}
	for i := 0; i < numberOfVerses; i++ {
		verses = append(verses, Verse{
			Ayah:        fmt.Sprintf("sura %v, verse %v", suraNumber, startVerse+i),
			SuraNumber:  suraNumber,
			VerseNumber: startVerse + i,
		})
	}
	return verses, nil
}
