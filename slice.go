package main

import "fmt"

func main() {
	names := [...]string{"Rohim", "Muhamad", "Cathryn", "Alice"}
	slice1 := names[0:2]
	fmt.Println(slice1)

	slice2 := names[:4]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"

	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Baru")
	daysSlice2[0] = "Sabtu Lama"

	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Rohim"
	newSlice[1] = "Muhamad"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Cathryn")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Dinda"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
}
