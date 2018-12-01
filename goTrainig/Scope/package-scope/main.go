package main

import (
	"fmt"

	"./vis"
)

var x int = 31

func main() {
	fmt.Println(x)
	someFunction()
	fmt.Println(vis.Y)
}

func someFunction() {
	fmt.Println(x)
}
