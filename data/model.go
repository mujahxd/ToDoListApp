package data

import (
	"database/sql"
)

type Model struct {
	conn *sql.DB
}

func (pm *Model) SetSQLConnection(db *sql.DB) {
	pm.conn = db
}

func (pm *Model) Login(username string, password string) (*User, error) {
	var user = User{}

	err := pm.conn.QueryRow("SELECT fullname, username FROM user where username = ? AND password = ?", username, password).Scan(&user.Fullname, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
