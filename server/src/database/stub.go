package database

import (
	"errors"
	"fmt"

	"github.com/cglotr/koran/server/src/quran"
)

// Stub .
type Stub struct{}

// GetVerses .
func (s *Stub) GetVerses(suraNumber, startVerse, numberOfVerses int) ([]quran.Verse, error) {
	if suraNumber < 1 {
		return nil, errors.New("invalid suraNumber")
	}
	if startVerse < 1 {
		return nil, errors.New("invalid startVerse")
	}
	if numberOfVerses < 1 {
		return nil, errors.New("invalid numberOfVerses")
	}
	verses := []quran.Verse{}
	for i := 0; i < numberOfVerses; i++ {
		verses = append(verses, quran.Verse{
			Ayah:        fmt.Sprintf("sura %v, verse %v", suraNumber, startVerse+i),
			SuraNumber:  suraNumber,
			VerseNumber: startVerse + i,
		})
	}
	return verses, nil
}
