package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	for {

		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please wait. Processing.....",c)
		time.Sleep(time.Duration(10)* time.Millisecond)
	}
}