package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	city := strings.ToLower(os.Args[1])

	switch city {
	case "paris", "lyon" :
		fmt.Println("France")
	case "tokyo" :
		fmt.Println("Japan")
	default :
		fmt.Println("Where?")
	}

	i := 10

	switch  {
	case i > 0 :
		fmt.Println("Positive")
	case i < 0 :
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	

	switch j := 142; {
	case j > 100 :
		fmt.Print("Big ")
		fallthrough
	case j > 0 :
		fmt.Print("Positive ")
		fallthrough
	default :
		fmt.Print("Number\n")
	}
}