package main

import (
	"fmt"
	"os"
)


func main(){

	if len(os.Args) != 2 {
		fmt.Println("Enter a month.")
		return
	}

	switch m := os.Args[1]; m {
	case "dec", "jan", "feb" :
		fmt.Println("Winter")
	case "mar", "apr", "may" :
		fmt.Println("Spring")
	case "jun", "jul", "aug" :
		fmt.Println("Summer")
	case "sep", "oct", "nov" :
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}