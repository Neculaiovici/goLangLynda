package main

import (
	"fmt"
)

func main() {
	a := 31
	fmt.Println("value of a -", a)
	fmt.Println("address of a -", &a)

	var b *int = &a
	fmt.Println("Values of pointer b: ")
	fmt.Println("value address of b -", b)
	fmt.Println("address in memory of b -", &b)
	fmt.Println("value of b -", *b) // * is dereferencing

}
