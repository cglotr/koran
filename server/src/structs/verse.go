package structs

// Verse .
type Verse struct {
	SuraNumber  int    `json:"sura_number,omitempty"`
	VerseNumber int    `json:"verse_number,omitempty"`
	Ayah        string `json:"ayah,omitempty"`
}
