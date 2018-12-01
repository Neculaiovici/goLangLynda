package main

import (
	"fmt"
)

func main() {
	// block level scope
	// uncomment, doesn't work
	// fmt.Println(y)
	y := 22
	fmt.Println(y)

	shadow := shadow(10)
	//shadow is the result not the function
	// don't do these, bad coding practice to shadow the variable
	fmt.Println(shadow)
}

func foo() {
	// uncomment, doesn't work
	// fmt.Println(y)
	t := 1
	fmt.Print(t)
}

// bad practice
func shadow(z int) int {
	return 28 + z
}
