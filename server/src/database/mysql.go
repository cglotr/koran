package database

import (
	"database/sql"

	"github.com/cglotr/koran/server/src/quran"
)

// Mysql .
type Mysql struct {
	Db *sql.DB
}

// CreateUser .
func (m *Mysql) CreateUser(uid string) (int, error) {
	result, err := m.Db.Exec(
		`
		INSERT INTO user (uid)
		VALUES (?)
		;
		`,
		uid,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// CreateUserQuran .
func (m *Mysql) CreateUserQuran(userID, suraID, verseID int) error {
	_, err := m.Db.Exec(
		`
		INSERT INTO user_quran (user_id, sura_id, verse_id)
		VALUES (?, ?, ?)
		;
		`,
		userID,
		suraID,
		verseID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserQuran .
func (m *Mysql) DeleteUserQuran(userID, suraID, verseID int) error {
	_, err := m.Db.Exec(
		`
		DELETE FROM user_quran
		WHERE
			user_id = ?
			AND
			sura_id = ?
			AND
			verse_id = ?
		;
		`,
		userID,
		suraID,
		verseID,
	)
	if err != nil {
		return err
	}
	return nil
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

// GetUser .
func (m *Mysql) GetUser(uid string) (*User, error) {
	rows, err := m.Db.Query(
		`
		SELECT id, token
		FROM user
		WHERE uid = ?
		;
		`,
		uid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id int
	var token string
	if rows.Next() {
		rows.Scan(&id, &token)
	}
	return &User{
		ID:    id,
		UID:   uid,
		Token: token,
	}, nil
}

// GetUserByID .
func (m *Mysql) GetUserByID(id int) (*User, error) {
	rows, err := m.Db.Query(
		`
		SELECT uid, token
		FROM user
		WHERE id = ?
		;
		`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var uid string
	var token string
	if rows.Next() {
		rows.Scan(&uid, &token)
	}
	return &User{
		ID:    id,
		UID:   uid,
		Token: token,
	}, nil
}

// GetVerses .
func (m *Mysql) GetVerses(suraNumber, startVerse, numberOfVerses int) ([]quran.Verse, error) {
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
	verses := []quran.Verse{}
	var id int
	var suraID int
	var verseID int
	var ayah string
	for rows.Next() {
		rows.Scan(&id, &suraID, &verseID, &ayah)
		verses = append(verses, quran.Verse{
			SuraNumber:  suraID,
			VerseNumber: verseID,
			Ayah:        ayah,
		})
	}
	return verses, nil
}

// IsUserExist .
func (m *Mysql) IsUserExist(uid string) bool {
	rows, err := m.Db.Query(
		`
		SELECT id, uid, token
		FROM user
		WHERE uid = ?
		;
		`,
		uid,
	)
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}

// UpdateUserToken .
func (m *Mysql) UpdateUserToken(id int, token string) bool {
	_, err := m.Db.Exec(
		`
		UPDATE user SET token = ? WHERE id = ?
		;
		`,
		token,
		id,
	)
	if err != nil {
		return false
	}
	return true
}
