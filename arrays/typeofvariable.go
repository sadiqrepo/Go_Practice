package main

import (
	"fmt"
	"reflect"
)

func main(){
	var intArray [5]int
	var strArray [5]int
	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())

}