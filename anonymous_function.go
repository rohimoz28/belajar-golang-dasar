package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//CARA 1
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("rohim", blacklist)

	//CARA 2
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
