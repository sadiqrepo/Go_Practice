package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	
	validops = "%*/+-"
	usage = "Usage: [op="+validops+"] [size]"
	missingSizeMsg = "Size is missing"
	wrongSizeMsg = "Wrong size"
	invalidOprMsg = `Invalid operator.
Valid ops one of: `+validops
	invalidOpr = -1

)

func main(){

	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println(missingSizeMsg)
		fallthrough
	case l < 1:
		fmt.Println(usage)
		return
	}

	op := args[0]
	if strings.IndexAny(op, validops) == invalidOpr{
		fmt.Println(invalidOprMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println(wrongSizeMsg)
		return
	}


	fmt.Printf("%5s", "X")
		for i:=0; i <= size; i++{
			fmt.Printf("%5d",i)
		}

		fmt.Println()
		for i:=0; i <= size; i++{
			fmt.Printf("%5d",i)

			for j:=0; j<=size; j++ {

				var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}
				fmt.Printf("%5d",res)
			}
			fmt.Println()
		}



	
}