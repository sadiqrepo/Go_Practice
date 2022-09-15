package main

import (
	"fmt"
)

func main(){

	var (
		blue = [3]int{6,9,3}
		red = [3]int{6,9,3}
		black = [...]int{6,9,3}
		white = [3]int{6,9}
	)

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)

	fmt.Println("Blue equal to red?", blue == red)
	fmt.Println("Black equal to red?", black == red)
	fmt.Println("Blue equal to white?", blue == white)

}