package main

import "fmt"

type udemy int

var x udemy
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("x Type: %T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("Value of x: %v\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("Value of y: %T\n", y)
}
