package main

import (
	"fmt"
	"reflect"
)


func multiVariadic(val ...interface{}){

	for _, v := range val{
		//fmt.Println(v, "--", reflect.TypeOf(v))
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main(){
	multiVariadic(1, "red", 25.5, []string{"foo", "bar", "buzz"}, map[string]int{"apple": 5, "Oranges": 10})
}