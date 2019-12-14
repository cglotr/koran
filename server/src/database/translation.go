package database

// Translation .
type Translation struct {
	SuraNumber  int    `json:"sura_number,omitempty"`
	VerseNumber int    `json:"verse_number,omitempty"`
	Translation string `json:"translation,omitempty"`
}
