package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){

	i := 10
	s := strconv.Itoa(i)
	fmt.Println(s, reflect.ValueOf(s).Kind())

}