package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	var (
		leap bool
		year int
		err error
	)
	if len(os.Args) !=2 {
		fmt.Println("Give me a year number")
		return
	} else if year, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n",os.Args[1])
		return
	} else if year % 400 == 0  {
		leap = true
	} else if year % 100 == 0{
		leap = false
	} else if year % 4 == 0 {
		leap = true
	}


	if leap { fmt.Printf("%d is a leap year.\n",year)
	} else {
		fmt.Printf("%d is not a leap year.\n",year)
	}
}