package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Rohim",
		"address": "Tangerang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(len(person))

	person["name"] = "Cathryn"
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Naruto"
	book["author"] = "Masashi Kishimoto"
	book["wrong"] = "Upps"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)

}
