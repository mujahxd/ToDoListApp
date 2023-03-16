package main

import (
	"fmt"
	"todolistapp/config"
	// "simpleSQL/produk"
)

// ini remark ugi
func main() {
	var koneksi = config.InitSQL()
	// var mdl = produk.ProdukModel{}
	// mdl.SetSQLConnection(koneksi)

	var menu int
	for menu != 99 {
		fmt.Println("Selamat datang di Aplikasi Todolist")
		fmt.Println("1.login")
		fmt.Println("2.Regiter")
		fmt.Println("3. Exit")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			var Username string
			var Password string
			fmt.Println("Silahkan Login")
		}
	}
	fmt.Println(koneksi)
}
