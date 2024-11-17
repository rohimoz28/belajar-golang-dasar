package main

import "fmt"

type Filter func(name string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("John", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
