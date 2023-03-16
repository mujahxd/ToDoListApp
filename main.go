package main

import (
	"fmt"
	"todolistapp/config"
	"todolistapp/data"
)

// ini remark ugi
func main() {
	var koneksi = config.InitSQL()
	var mdl = data.Model{}
	mdl.SetSQLConnection(koneksi)

	var menu int
	for menu != 99 {
		fmt.Println("Selamat Datang di Aplikasi ToDo List")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			var username string
			var password string
			fmt.Println("Silahkan Login")
			fmt.Println("Masukan Username")
			fmt.Scanln(&username)
			fmt.Println("Masukan Password")
			fmt.Scanln(&password)
			res, err := mdl.Login(username, password)
			if err != nil {
				fmt.Println("password/username salah")
				break
			}
			fmt.Println("halo selamat datang " + res.Fullname)
			menu = 99

		case 2:
			fmt.Println("halo selamat datang, apa yang akan anda lakukan hari ini")

		}
	}

}

// 	"log"
// 	"todolistapp/config"
// 	"todolistapp/user"
// )

// func main() {
// 	db, err := config.ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var user = user.UserModel{}
// 	user.SetSQLConnection(db)

// 	var menu int
// 	for menu != 99 {
// 		fmt.Println("Selamat Datang di Aplikasi ToDo List")
// 		fmt.Println("1. Login")
// 		fmt.Println("2. Register")
// 		fmt.Println("0. Exit")
// 		fmt.Print("Pilih menu: ")
// 		fmt.Scanln(&menu)

// 	}

// }
