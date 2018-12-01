package main

import (
	"fmt"

	rever "../packageUtil"
)

func main() {
	simple := "Best goLang !oOG"
	rvs := "drow esrever"
	fmt.Println(simple)
	fmt.Println(rever.Reverse(rvs))
	fmt.Println(rvs)
	fmt.Println(rever.Reverse(simple))
}
