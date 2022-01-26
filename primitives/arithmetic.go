package main

import (
	"fmt"
)

func main(){

	a := 10
	b := 6

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println( a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

	fmt.Println(a << 3)
	fmt.Println( a >> 3)


	var f float64 = 3.14
	fmt.Printf("%v, %T\n", f, f)

	var s string
	s = "This is a String"
	fmt.Printf("%v, %T\n", s, s)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
	


}