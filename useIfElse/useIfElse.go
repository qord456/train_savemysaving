package useifelse

import "fmt"

func Useifelse() {
	var nama string
	var peran string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	for peran == "" {
		fmt.Print("Masukkan peran (superhero/monster): ")
		fmt.Scan(&peran)
	}

	if nama == "" {
		fmt.Println("Nama Harus Di Isi")
	} else if peran == "superhero" {
		fmt.Println("Selamat Datang Superhero", nama+", Kalahkan Semua Monster Di Muka Bumi")
	} else if peran == "monster" {
		fmt.Println("Selamat Datang Monster", nama+", Hancurkan Semua Superhero Yang Ada")
	} else {
		fmt.Println("Selamat Datang", nama+", Pilih Peranmu Untuk Melanjutkan Game Ini")
		Useifelse()
	}
}
