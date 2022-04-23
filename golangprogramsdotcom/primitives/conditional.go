package main

import (
	"fmt"
)

func main() {

	check1 := 1 == 1
	check2 := 1 == 2

	fmt.Printf("%v, %T\n", check1, check1)
	fmt.Printf("%v, %T\n", check2, check2)
}
