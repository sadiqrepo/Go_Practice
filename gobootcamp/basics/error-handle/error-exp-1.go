package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	s := strconv.Itoa(42)
	fmt.Println(s)

	age := os.Args[1]

	n, err := strconv.Atoi(age)
	fmt.Println("Converted number: ", n)
	//fmt.Println("Returned error value: ",err)

	if err != nil {
		fmt.Println("ERROR: ",err)
		return
	}

	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}