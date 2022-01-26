package main

import (
	"fmt"
	"reflect"
)


func main(){

	var a = 10
	var b = "Lisa"

	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(b).Kind())
}