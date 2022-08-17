package main

import (
	"fmt"
)

func main(){

	name := "Bryan"
	lastname := "Adams"

	fmt.Printf("My name is %s and my last name is %s.\n", name, lastname)

	msg := "My name is %s and my last name is %s.\n"
	fmt.Printf(msg, name, lastname)
}