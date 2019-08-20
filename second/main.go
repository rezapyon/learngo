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

	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
}
