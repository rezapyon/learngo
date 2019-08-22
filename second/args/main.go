package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("%#v\n", os.Args)
	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("1st Argument:", os.Args[1])
	// fmt.Println("2nd Argument:", os.Args[2])
	// fmt.Println("3rd Argument:", os.Args[3])

	// fmt.Println("Number of items inside os.Args:", len(os.Args))
	// $ go run main.go chaeyoung is beautiful

	name := os.Args[1]
	fmt.Println("Hello great", name, "!")

	name, age := "Gandalf", 2019
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")

}
