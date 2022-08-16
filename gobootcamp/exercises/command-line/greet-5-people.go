package main

import (
	"fmt"
	"os"
)

func main(){
	
	n := len(os.Args)-1
	person1 := os.Args[1]
	person2 := os.Args[2]
	person3 := os.Args[3]
	person4 := os.Args[4]
	person5 := os.Args[5]

	fmt.Println("There are", n, "people!")
	fmt.Println("Hello great", person1, "!")
	fmt.Println("Hello great", person2, "!")
	fmt.Println("Hello great", person3, "!")
	fmt.Println("Hello great", person4, "!")
	fmt.Println("Hello great", person5, "!")
	fmt.Println("Nice to meet you all.")
}