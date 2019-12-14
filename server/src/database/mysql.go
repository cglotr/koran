package database

import (
	"database/sql"
)

// Mysql .
type Mysql struct {
	Db *sql.DB
}

// GetTranslations .
func (m *Mysql) GetTranslations(name string, suraNumber, startVerse, numberOfVerses int) ([]Translation, error) {
	rows, err := m.Db.Query(
		`
		SELECT id, sura_id, verse_id, translation
		FROM translation
		WHERE
			name = ?
			AND
			sura_id = ?
			AND
			verse_id >= ?
			AND
			verse_id < ?
		;
		`,
		name,
		suraNumber,
		startVerse,
		startVerse+numberOfVerses,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	translations := []Translation{}
	var id int
	var suraID int
	var verseID int
	var translation string
	for rows.Next() {
		rows.Scan(&id, &suraID, &verseID, &translation)
		translations = append(translations, Translation{
			SuraNumber:  suraID,
			VerseNumber: verseID,
			Translation: translation,
		})
	}
	return translations, nil
}

// GetVerses .
func (m *Mysql) GetVerses(suraNumber, startVerse, numberOfVerses int) ([]Verse, error) {
	rows, err := m.Db.Query(
		`
		SELECT id, sura_id, verse_id, ayah
		FROM quran
		WHERE
			sura_id = ?
			AND
			verse_id >= ?
			AND
			verse_id < ?
		;
		`,
		suraNumber,
		startVerse,
		startVerse+numberOfVerses,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	verses := []Verse{}
	var id int
	var suraID int
	var verseID int
	var ayah string
	for rows.Next() {
		rows.Scan(&id, &suraID, &verseID, &ayah)
		verses = append(verses, Verse{
			SuraNumber:  suraID,
			VerseNumber: verseID,
			Ayah:        ayah,
		})
	}
	return verses, nil
}
