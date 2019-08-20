package main

import "fmt"

func main() {
	fmt.Println(-200, -100, 0, 50, 100)       //int
	fmt.Println(-50.5, 20.5, 0., 1.0, 100.56) //float64
	fmt.Println(true, false)                  //bool
	fmt.Println(
		"", //empty string
		"hi",
		"안녕 하세요",
	)

	// var speed int
	var (
		speed int
		heat  float64

		off   bool
		brand string
	)

	var speeds, velocity int

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
	fmt.Println(speeds, velocity)

	//=============== inference =====================
	fmt.Println("============= inference ==============")

	var safe bool = true
	var safee = true //go automatically deduce that the variable is bool type
	safeee := true   // short declaration statement

	fmt.Println(safe, safee, safeee)

	name := "Carl" //short declaration can not be used on package scope, use var name = "Carl" instead
	isScientist := true
	age := 62
	degree := 5.

	fmt.Println(name, isScientist, age, degree)

	// multiple short declaration
	light, room := true, 4

	fmt.Println(light, room)
}
