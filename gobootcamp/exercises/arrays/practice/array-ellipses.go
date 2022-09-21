package main

import (
	"fmt"
	"reflect"
)

func main(){

	elp := [...]int{10, 20, 30, 40, 50}
	fmt.Println(elp)
	fmt.Println(reflect.ValueOf(elp).Kind())
	fmt.Println(len(elp))
}