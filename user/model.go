package user

import "database/sql"

type UserModel struct {
	conn *sql.DB
}

func (um *UserModel) SetSQLConnection(db *sql.DB) {
	um.conn = db
}
