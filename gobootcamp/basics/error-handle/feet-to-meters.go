package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

func main(){

	if len(os.Args) > 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	feet := os.Args[1]

	ft, err := strconv.ParseFloat(feet, 64)

	if err != nil{
		fmt.Printf("error: '%s' is not a number.\n", feet)
		return
	}

	meters := ft * 0.3048

	fmt.Printf("%g feet is %g meters.\n", ft, meters)
}