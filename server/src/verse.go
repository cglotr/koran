package main

type verse struct {
	SuraNumber  int    `json:"sura_number,omitempty"`
	VerseNumber int    `json:"verse_number,omitempty"`
	Ayah        string `json:"ayah,omitempty"`
}
