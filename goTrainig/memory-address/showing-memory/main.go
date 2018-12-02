package main

import (
	"fmt"
)

func main() {
	a := 31
	fmt.Println("value of stored a -", a)
	fmt.Println("showing a memory address -", &a)
	fmt.Printf("address: %d  \n", &a)
}
