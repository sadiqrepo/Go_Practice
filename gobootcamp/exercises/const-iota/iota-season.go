package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func main() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		_ = iota * 3
		Spring 
		Summer 
		Fall   
		Winter 
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}