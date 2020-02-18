package quran

import "github.com/cglotr/koran/server/src/structs"

// VersesGetter .
type VersesGetter interface {
	GetVerses(suraNumber, startVerse, numberOfVerses int) ([]structs.Verse, error)
}
