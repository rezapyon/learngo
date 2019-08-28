package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("sum :", 3+2)    //sum int
	fmt.Println("sum :", 2+3.5)  //sum float64
	fmt.Println("dif :", 3-1)    //difference int
	fmt.Println("dif :", 3-0.5)  //difference float64
	fmt.Println("prod :", 4*5)   //product int
	fmt.Println("prod :", 5*2.5) //product float64
	fmt.Println("quot :", 8/2)   //quotient intgit
	fmt.Println("quot :", 8/1.5) //quotient float64

	fmt.Println("rem :", 8%3) //remainder only for ints
	// fmt.Println("rem :", 8.0%3)   //remainder error

	f := 8.0
	fmt.Println("rem :", int(f)%3)

	var ratio float64 = 3 / 2
	fmt.Println(ratio) //result is 1

	ratio = float64(3) / 2
	fmt.Println(ratio) //result is 1.5

	celcius := 35.
	fahrenheit := (9*celcius + 160) / 5

	fmt.Printf("%g C is %g F\n", celcius, fahrenheit)

	name := "Reza"
	fmt.Println(len(name))

	name = "손채영"
	fmt.Println(len(name)) // false count
	fmt.Println(
		utf8.RuneCountInString(name))

}
