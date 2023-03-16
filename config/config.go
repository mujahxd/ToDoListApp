package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:lazuardi12@tcp(127.0.0.1:3306)/todolistapp")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if db.Ping() != nil {
		fmt.Println(db.Ping().Error())
		return nil
	}

	return db
}

// 	"log"

// 	_ "github.com/lib/pq"
// )

// func ConnectDB() (*sql.DB, error) {
// 	// konfigurasi koneksi
// 	host := "localhost"
// 	port := "5432"
// 	user := "postgres"
// 	password := "s3cur1tyIT"
// 	dbname := "todolistapp"

// 	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	// buat koneksi

// 	db, err := sql.Open("postgres", dbinfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// tes koneksi
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db, nil

// }
