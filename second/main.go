package main

import (
	"fmt"
	"path"
)

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
	var light bool
	light, room := true, 4 //assign value to light variable and initialize new variable called room

	fmt.Println(light, room)

	pathseparator("css/main.css")
	convertionexample()
}

func pathseparator(str string) {
	dir, file := path.Split(str)
	// _, file := path.Split(str) //if you want to use file variable only

	fmt.Println("Dir: ", dir)
	fmt.Println("File: ", file)
}

func convertionexample() {
	speed := 100
	force := 2.5

	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
