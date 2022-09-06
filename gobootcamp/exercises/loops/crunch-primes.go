package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := os.Args[1:]

	for _, v := range input {
		a, err := strconv.Atoi(v)

		if err != nil {
			continue
		}

		var flag bool

		if a <= 1 {
			flag = false
		} else if a > 1 {

			for i := 2; i < a/2; i++ {
				if a%i == 0 {
					flag = false
					break
				}
			}
		} else {
			flag = true
		}
		if flag == true {
			fmt.Printf("%d ", a)
		}
	}
}
