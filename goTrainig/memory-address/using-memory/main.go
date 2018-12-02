package main

import (
	"fmt"
)

const hectar float64 = 10000

func main() {
	var mp2 float64
	fmt.Print("Conversie metru patrat in hectar \nAdaugati valoarea: ")
	fmt.Scan(&mp2)
	ha := mp2 / hectar
	fmt.Println(mp2, " metri patrati este ", ha, " ha.")
}
