package main

import (
	"fmt"
	"os"
)

func main(){

	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%q\n",v)
	}

	for _, v := range os.Args[1:] {	
		fmt.Printf("%q ",v)
	}
	fmt.Println()
}