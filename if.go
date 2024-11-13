package main

import "fmt"

func main() {
	name := "Cathryn"

	if name == "Rohim" {
		fmt.Println("Saya Rohim")
	} else if name == "Cathryn" {
		fmt.Println("Hai Cathryn")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	//IF SHORT STATEMENT
	if length := len(name); length > 9 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
