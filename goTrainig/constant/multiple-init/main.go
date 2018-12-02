package main

import (
	"fmt"
)

// const (
// 	A = iota
// 	B = iota
// 	C = iota
// )

const (
	A = iota
	B
	C
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
