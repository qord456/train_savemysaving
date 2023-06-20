package main

import (
	useifelse "challenge-git-team/useIfElse"
	"fmt"
)

func main() {
	var menu int
	fmt.Println("1. Bermain Dengan Perulangan")
	fmt.Println("2. Bermain Dengan Irisan Array")
	fmt.Println("3. Menggunakan Switch Case")
	fmt.Println("4. Menggunakan If Else")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		// Ubah ini untuk packagemu
	case 2:
		// Ubah ini untuk packagemu
	case 3:
		// Ubah ini untuk packagemu
	case 4:
		useifelse.Useifelse()
	}

}