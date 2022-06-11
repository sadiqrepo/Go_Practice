package main

import(
	"fmt"
)

func main(){

	fmt.Print("Enter a number: ")
	var num int
	fmt.Scan(&num)

	switch {
	case num%2 == 0:
		fmt.Println("EVEN NUMBER")
	case num%2 !=0:
		fmt.Println("ODD NUMBER")
	default:
		fmt.Println("Enter VALID Number!!")
	}
}