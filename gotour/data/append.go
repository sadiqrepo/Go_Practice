package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main(){

	var a []int
	printSlice(a)

	a = append(a, 0)
	printSlice(a)

	a = append(a, 1)
	printSlice(a)

	a = append(a, 2)
	printSlice(a)

	a = append(a, 3)
	printSlice(a)

	a = append(a, 4)
	printSlice(a)

	a = append(a, 5)
	printSlice(a)

	a = append(a, 6)
	printSlice(a)

	a = append(a, 7)
	printSlice(a)

	a = append(a, 8)
	printSlice(a)
}