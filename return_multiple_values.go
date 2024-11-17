package main

import "fmt"

func getFullName() (string, string) {
	return "Rohim", "Muhamad"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//Ignore Value
	firstName, _ = getFullName()
	fmt.Println(firstName)
}
