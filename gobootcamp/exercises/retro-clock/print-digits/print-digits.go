package main

import (
	"fmt"
)

func main(){

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}


	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}


	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}


/*	for line := 0 ; line < 5; line++ {
		
		for digit := range digits {
			fmt.Print(digits[digit][line]," ")
		}
		fmt.Println()
	}

	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line]," ")
		}
		fmt.Println()
	}*/

	for _, digit := range digits {
		for _, line := range digit {
			fmt.Println(line)
		}
		fmt.Println()
	}



}