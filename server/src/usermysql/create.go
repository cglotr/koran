package usermysql

// Create .
func (m *Mysql) Create(uid string) (int, error) {
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

// Exist .
func (m *Mysql) Exist(uid string) bool {
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
