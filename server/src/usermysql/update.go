package usermysql

// UpdateToken .
func (m *Mysql) UpdateToken(id int, token string) bool {
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
