package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rohim"
	names[1] = "Muhamad"
	names[2] = "Nur"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 80}

	var fruits = [...]string{"apple", "manggo", "grape"} // definisi array tanpa deklarasi jumlah array di awal

	fmt.Println(values)
	fmt.Println(len(fruits))
}
