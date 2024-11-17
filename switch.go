package main

import "fmt"

func main() {

	name := "Rohim"

	switch name {
	case "Rohim":
		fmt.Println("I'm Rohim")
	case "Cathryn":
		fmt.Println("I'm Cathryn")
	default:
		fmt.Println("Boleh Kenalan Gak?")
	}

	//SHORT STATEMENT

	switch length := len(name); length < 6 {
	case true:
		fmt.Println("Nama Sudah Benar")
	case false:
		fmt.Println("Nama Terlalu Panjang")
	}
}
