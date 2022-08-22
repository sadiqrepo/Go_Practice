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

	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || arg1 >= arg2 {
		fmt.Println("Invalid input")
		return
	}

	var (
		sum int
		i int = arg1

	)

	for  {
		if i > arg2 {
			break
		} else if i%2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Print(i)
		if i < arg2 - 1 {
				fmt.Print(" + ")
			}
		i++
	}
	fmt.Printf(" = %d\n",sum)
}