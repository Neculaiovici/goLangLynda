package main

import (
	"fmt"
)

const b string = "global constant"
const a = 28

func main() {
	const a = 42
	const b string = "inside main constant"

	fmt.Println("constant b - ", b)
	fmt.Println("constant a - ", a)
}
