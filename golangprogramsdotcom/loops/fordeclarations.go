package main

import (
	"fmt"
)

func main(){

	k := 1
	for ; k <=10; k++ {
		fmt.Printf("%v",k)
	}
	fmt.Println()

	k = 1
	for k <=5 {
		fmt.Printf("%v",k)
		k++
	}
	fmt.Println()

	for k := 2; ;k++{
		if k > 5{
			break
		}
		fmt.Printf("%v",k)
	}
	fmt.Println()

	for k:=6; k <=10; k++ {
		fmt.Printf("%v",k)
	}
	fmt.Println()
	
}