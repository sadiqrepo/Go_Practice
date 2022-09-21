package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	args := os.Args[1:]

	switch l := len(args); {

	case l == 0:
		fmt.Println("Please give me upto 5 numbers.")
		return
	case l > 5:
		fmt.Println("Sorry. Go arrays are fixed.",
						"So, for now, I'm only supporting sorting 5 numbers...")
						return

	}

	var nums [5]float64

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		nums[i] = n
	}

	for range nums {

		for i, v := range nums {

			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		

	}

	

	fmt.Println(nums)
}