package main

import "fmt"

func main(){

	var names = [4]string{
		"John",
		"Elton",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	var a = names[0:2]
	var b = names[1:3]
	fmt.Println(a)
	fmt.Println(b)

	b[0] = "XXX"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(names)


	
}