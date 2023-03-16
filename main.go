package main

import (
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

	var user = user.UserModel{}
	user.SetSQLConnection(db)

	var menu int
	for menu != 99 {
		fmt.Println("Selamat Datang di Aplikasi ToDo List")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)

	}

}
