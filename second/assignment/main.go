package main

import (
	"fmt"
	"time"
)

func main() {
	//=========================SIngle Assignment=====================================//
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

	//=========================Multiple Assignment=====================================//

	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	//swapping variables

	var (
		speeds    = 100
		prevSpeed = 50
	)

	speeds, prevSpeed = prevSpeed, speeds

	fmt.Println(speeds, prevSpeed)
}
