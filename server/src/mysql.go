package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db *sql.DB
}

func (m *mysql) getVerses(suraNumber, startVerse, numberOfVerses int) ([]verse, error) {
	rows, err := m.db.Query(
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
	verses := []verse{}

	var id int
	var suraID int
	var verseID int
	var ayah string

	for rows.Next() {
		rows.Scan(&id, &suraID, &verseID, &ayah)
		verses = append(verses, verse{
			SuraNumber:  suraID,
			VerseNumber: verseID,
			Ayah:        ayah,
		})
	}

	return verses, nil
}
