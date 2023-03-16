package data

import (
	"database/sql"
)

type model struct {
	conn *sql.DB
}

func (pm *model) SetSQLConnection(db *sql.DB) {
	pm.conn = db
}

func (pm model) Login([]User, error) {
	rows, err := pm.conn.Query("SELECT barcode, nama, qty, harga, input_oleh FROM produk")
}
