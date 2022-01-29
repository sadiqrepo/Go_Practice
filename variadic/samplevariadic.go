package main

import (
	"fmt"
)

func variadicsample(s ...string){
	//fmt.Println(s[0])
	//fmt.Println(s[2])
	fmt.Println(s)
}

func calculation(str string, a ...int) int{
	
	area := 1

	for _, val := range a {
		if str == "Rectangle"{
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func main(){

	variadicsample()
	variadicsample("Red", "Blue", "Green", "Yellow")

	fmt.Println(calculation("Rectangle", 20 , 30))
	fmt.Println(calculation("Square", 20 ))
}