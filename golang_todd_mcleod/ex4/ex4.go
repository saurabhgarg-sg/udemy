package main

import "fmt"

type udemy int

var x udemy

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%v\n", x)
}
