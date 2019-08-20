package main

import "fmt"

func main() {
	var safe bool = true
	var safee = true //go automatically deduce that the variable is bool type
	safeee := true   // short declaration statement

	fmt.Println(safe, safee, safeee)

	name := "Carl"
	isScientist := true
	age := 62
	degree := 5.

	fmt.Println(name, isScientist, age, degree)
}
