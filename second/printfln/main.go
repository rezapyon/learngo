package main

import "fmt"

func main() {
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Planet: %s\n", planet)
	// fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Distance: %d millions kms\n", distance)
	// fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	// fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)
	// fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)
	fmt.Printf("%s is %d away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)

	//formating float number
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	fmt.Printf("Orbital Period: %.5f days\n", orbital)
}
