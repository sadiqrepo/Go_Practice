package main

import "fmt"

func main(){

	var sum int

	for i:=1; i<=10; i++ {
		fmt.Printf("%d",i)
		sum += i
		if sum < 55 {
			fmt.Print(" + ")
		}
	}

	fmt.Println("=",sum)
}