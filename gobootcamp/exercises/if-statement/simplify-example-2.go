package main

import "fmt"


func main() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool

	if radius >= 50 {
		if radius >= 100 {
			if radius >= 200 {
				big = true
			}
		}
	}

	if big != true {
		fmt.Println("I don't know.")
	} else if !(isSphere == false) {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}