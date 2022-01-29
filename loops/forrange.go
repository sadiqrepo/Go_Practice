package main

import (
	"fmt"
)

func main(){

	strDist := map[string]string {"Japan" : "Tokyo", "Canada" : "Vancouver", "India" : "New Delhi", "Germany" : "Berlin"}
	for index, element := range strDist{
		fmt.Println("Index: ", index, "Element: ", element)
	}


	for key := range strDist{
		fmt.Println(key)
	}

	for _, value := range strDist{
		fmt.Println(value)
	}

	for range "Hello"{
		fmt.Println("Hello")
	}

	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}