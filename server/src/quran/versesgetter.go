package quran

// VersesGetter .
type VersesGetter interface {
	GetVerses(suraNumber, startVerse, numberOfVerses int) ([]Verse, error)
}

// Verse .
type Verse struct {
	SuraNumber  int    `json:"sura_number,omitempty"`
	VerseNumber int    `json:"verse_number,omitempty"`
	Ayah        string `json:"ayah,omitempty"`
}
