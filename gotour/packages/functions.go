package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y 
}

func main(){
	fmt.Println("Sum of two numbers is", add(48, 2))
}