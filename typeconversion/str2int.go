package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main(){

	strVar := "100"
	if intVar, err := strconv.Atoi(strVar); err == nil{
		fmt.Println(intVar, reflect.TypeOf(intVar))
	}
	

}