package main

import (
	"fmt"
)

func main(){

	var (
		sum int
		j = 1
	) 

	for i := 0 ; i <= 1000; i++ {
		sum += i
	}

	fmt.Println(sum)

	for {
		if j > 5 { break }

		if j % 2 != 0 { 
			j++
			continue }
		sum += j
		j++
	}
	fmt.Println(sum)
}