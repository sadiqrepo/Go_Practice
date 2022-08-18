package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	} else if num, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a number.\n", os.Args[1])
		return
	} else if num % 8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8.\n",num)
	} else if num % 2 == 0 {
		fmt.Printf("%d is an even number.\n",num)
	} else {
		fmt.Printf("%d is an odd number.\n",num)
	}
}