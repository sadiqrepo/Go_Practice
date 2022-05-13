package main

import "fmt"

func main(){
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("Value of i is: ", i)
		sum += i
	}
	fmt.Println("Value of sum is: ", sum)
}