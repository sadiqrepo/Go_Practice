package main

import ("fmt")

func main(){

	intArray := [5]int{1,2,3,4,5}

	for index := 0; index < len(intArray); index++{
		fmt.Println(intArray[index])
	}
}