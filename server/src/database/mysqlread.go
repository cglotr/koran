package database

// ReadUserQuran .
func (m *Mysql) ReadUserQuran(userID, suraID, verseID int) (bool, error) {
	rows, err := m.Db.Query(
		`
		SELECT *
		FROM user_quran
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
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}
