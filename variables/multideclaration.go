package main


import (
	"fmt"
)


func main(){

	var fname, lname string = "John", "Whyte"
	fmt.Println(fname, lname)

	a, b, c := 10, 20, 30
	fmt.Println(a + b + c)

	item, price := "Mobile", 10000
	fmt.Println(item, "-", price)
}