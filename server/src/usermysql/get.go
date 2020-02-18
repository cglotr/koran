package usermysql

import "github.com/cglotr/koran/server/src/structs"

// Get .
func (m *Mysql) Get(uid string) (*structs.User, error) {
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
	return &structs.User{
		ID:    id,
		UID:   uid,
		Token: token,
	}, nil
}

// GetByID .
func (m *Mysql) GetByID(id int) (*structs.User, error) {
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
	return &structs.User{
		ID:    id,
		UID:   uid,
		Token: token,
	}, nil
}
