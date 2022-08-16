package main

import "fmt"

func main(){
	var brand string
	fmt.Printf("%q\n",brand)

	var name = "Google"

	fmt.Printf("%q\n", name)
	fmt.Printf("%s\n", name)

	fmt.Println("Hi\nHi")
	fmt.Println("Hi\\n\"Hi\"")
}