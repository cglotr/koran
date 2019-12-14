package database

import (
	"database/sql"
)

// Mysql .
type Mysql struct {
	Db *sql.DB
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
