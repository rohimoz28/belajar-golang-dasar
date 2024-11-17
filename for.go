package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan Ke-", counter)
	//	counter++
	//}
	//
	//fmt.Println("Selesai")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke-", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Rohim", "Muhamad", "Nur"}

	//Perulangan Manual
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	//Menggunakan For Range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	//Jika Tidak Butuh Index/Key
	for _, name := range names {
		fmt.Println(name)
	}
}
