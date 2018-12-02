package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("For loop with condition")
	x := 0
	for x < 10 {
		x++
		if x%2 != 0 {
			fmt.Println("Bum")
		} else if x%3 != 0 {
			fmt.Println("Bam")
		} else if x%3 != 0 || x%5 != 0 {
			fmt.Println("Cichi bum")
		}
	}

	y := time.Now()
	fmt.Println(y)
}
