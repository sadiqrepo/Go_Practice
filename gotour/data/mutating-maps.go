package main

import (
	"fmt"
)

func main(){

	var w = make(map[string]int)
	w["answer"] = 20
	fmt.Println(w)

	w["answer"] = 40
	fmt.Println(w)

	delete(w, "answer")
	fmt.Println(w)

	elem, ok := w["answer"]
	fmt.Println("The value: ", elem, "Is present?", ok)



}