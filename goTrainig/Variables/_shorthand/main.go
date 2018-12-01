package main

import (
	"fmt"
)

func main() {
	a := 100
	b := "String"
	c := 3.22
	d := false

	//printing variables
	// godoc.org/fmt format printing %s %T %v ...
	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)

	// declared 0 value, the default value
	var e int
	var f string
	var g float64
	var h bool
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
}
