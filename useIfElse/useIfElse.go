package useifelse

import "fmt"

func useifelse() {
	var nama string
	var peran string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan peran (superhero/monster): ")
	fmt.Scan(&peran)

	if nama == "" || peran == "" {
		fmt.Println("Nama dan Peran Harus Di Isi")
	} else if peran == "superhero" {
		fmt.Println("Selamat Datang Superhero", nama+", Kalahkan Semua Monster Di Muka Bumi")
	} else if peran == "monster" {
		fmt.Println("Selamat Datang Monster", nama+", Hancurkan Semua Superhero Yang Ada")
	} else {
		fmt.Println("Selamat Datang", nama+", Pilih Peranmu Untuk Melanjutkan Game Ini")
	}
}