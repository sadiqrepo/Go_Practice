package main

import (
	"fmt"
)

func main(){

	blue := [5]int{6, 9, 3, 2, 1}
	red := [5]int{6, 9, 3, 2, 1}

	fmt.Println("Are they equal? ")

	if blue == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red: %#v\n", red)


	type (
		bookcase [5]int
		cabinet [5]int
	)

	green := bookcase{6, 9, 3, 2, 1}

	if green == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("green: %#v\n", green)
	fmt.Printf("red: %#v\n", red)

	yellow := cabinet{6, 9, 3, 2, 1}

	if green == bookcase(yellow ){
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("green: %#v\n", green)
	fmt.Printf("yellow: %#v\n", yellow)

	gadgets := [3]string{"Confused Drone"}
    fmt.Printf("%q\n", gadgets)

	
	 


}