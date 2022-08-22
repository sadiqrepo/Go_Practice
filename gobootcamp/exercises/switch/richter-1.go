package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	
	/* if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	} */

	switch scale, _ := strconv.ParseFloat(os.Args[1], 64); {
	case scale >= 11 :
		fmt.Printf("%.2f is massive.\n", scale)
	case scale >= 8 :
		fmt.Printf("%.2f is great.\n", scale)
	case scale >= 7 :
		fmt.Printf("%.2f is major.\n", scale)
	case scale >= 6 :
		fmt.Printf("%.2f is strong.\n", scale)
	case scale >= 5 :
		fmt.Printf("%.2f is moderate.\n", scale)
	case scale >= 4 :
		fmt.Printf("%.2f is light.\n", scale)
	case scale >= 3 :
		fmt.Printf("%.2f is minor.\n", scale)
	case scale >= 2 :
		fmt.Printf("%.2f is very minor.\n", scale)
	case scale > 0 :
		fmt.Printf("%.2f is micro.\n", scale)
	default :
		fmt.Println("I couldn't get that, sorry.")
	}
}