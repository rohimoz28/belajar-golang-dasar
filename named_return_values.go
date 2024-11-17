package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rohim"
	middleName = "Muhamad"
	lastName = "Nur"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
