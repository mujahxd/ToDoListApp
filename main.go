package main

import (
	"database/sql"
	"fmt"
	"log"
	"todolistapp/config"
	"todolistapp/user"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var userModel = user.UserModel{}
	var user = user.User{}
	userModel.SetSQLConnection(db)

	var menu int
outerloop:
	for menu != 99 {
		fmt.Println("Selamat Datang di Aplikasi ToDo List")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
		switch menu {
		// fitur login
		case 1:
		// fitur register
		case 2:
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&user.Username)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&user.Password)
			fmt.Print("Masukkan Fullname: ")
			fmt.Scanln(&user.FullName)
			err := register(db, user.Username, user.Password, user.Password)
			if err != nil {
				fmt.Printf("Username sudah ada, GAGAL membuat Akun\n\n")

			} else {

				fmt.Printf("User BERHASIL dibuat!\n\n")
			}

		case 99:
			break outerloop
		}
	}

}

func register(db *sql.DB, username, password, fullname string) error {
	sqlStatement := "INSERT INTO public.\"User\" (username, password, full_name) VALUES ($1, $2, $3)"
	_, err := db.Exec(sqlStatement, username, password, fullname)
	if err != nil {
		return err
	}

	return nil
}
