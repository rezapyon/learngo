package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("Previous Name: ", prevName)
	fmt.Println("Current Name: ", name)
}
