package main

import (
	"fmt"
)

func main() {
	// block level scope
	// uncomment, dont work
	// fmt.Println(y)
	y := 22
	fmt.Println(y)
}

func foo() {
	// uncomment, dont work
	// fmt.Println(y)
	t := 1
	fmt.Print(t)
}
