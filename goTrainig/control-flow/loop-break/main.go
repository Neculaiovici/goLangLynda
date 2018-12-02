package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("Loop break")
	fmt.Println("Input number for loop")
	fmt.Scan(&x)
	for {
		x++
		fmt.Println(x)
		if x >= 10 {
			fmt.Println("End of loop")
			break
		} else {
			fmt.Println(x)
		}
	}
	fmt.Println("Continue condition loop")
	y := 0
	for {
		y++
		if y%2 == 0 {
			continue
		}
		fmt.Println(y)
		if y >= 5 {
			break
		}
	}
}
