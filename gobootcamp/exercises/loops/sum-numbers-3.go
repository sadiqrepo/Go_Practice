package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	if len(os.Args) != 3 {
		fmt.Println("Enter min & max")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil  {
		fmt.Println("Incorrect min value")
		return
	} else if err2 != nil {
		fmt.Println("Incorrect max value")
		return
	} else if min >= max {
		fmt.Println("Min cannot be greater than max value")
		return
	}

	var sum int 

	for i:= min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n",sum)
	
}