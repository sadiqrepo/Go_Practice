package main

import (
	"fmt"
	"reflect"
)

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice{
		panic("Invalid data type")
	}

	for i:=0; i < s.Len(); i++{
		if s.Index(i).Interface() == item{
			return true
		}
	}

	return false
}

func main(){

	iSlice1 := []string{"India", "US", "UK"}
	iSlice2 := []string{"Germany", "Canada"}

	iSlice1 = append(iSlice1, iSlice2...)

	fmt.Println(iSlice1)
	fmt.Println(itemExists(iSlice1, "India"))
}