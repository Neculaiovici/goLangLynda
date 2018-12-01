package main

import (
	"fmt"

	rever "../packageUtil"
)

func main() {
	simple := "Best goLang !oOG"
	rvs := "drow esrever"
	// Printing my variables from outside
	fmt.Println(rever.MyName)

	//Printing the functions
	fmt.Println(simple)
	fmt.Println(rever.Reverse(rvs))
	fmt.Println(rvs)
	fmt.Println(rever.Reverse(simple))
}
