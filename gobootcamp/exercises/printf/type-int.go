package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func main() {
	var s int = 3
	fmt.Printf("Type of %d is %[1]T\n",s)
}