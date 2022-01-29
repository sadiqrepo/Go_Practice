package main

import (
	"fmt"
)

func update(age *int, text *string){
	*age = *age + 5
	*text = *text + " Doe"
	return
}


func main(){
	
	var age int = 20
	var text string = "John"

	fmt.Println("Before: ", text, age)

	update(&age, &text)

	fmt.Println("After: ", text, age)
}