package database

// VersesGetter .
type VersesGetter interface {
	GetVerses(suraNumber, startVerse, numberOfVerses int) ([]Verse, error)
}
