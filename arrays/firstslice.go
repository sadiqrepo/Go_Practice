package main

import (
	"fmt"
)

func main() {

	a := []int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(a)
	//fmt.Printf("Length: %v\n", len(a))
	//fmt.Printf("Capacity: %v\n", cap(a))
	b := a[:]
	c :=a[3:]
	d := a[:6]
	e := a[3:6]
	f := a[0:10]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)


	a1 := make([]int, 3, 100)
	fmt.Println(a1)
	fmt.Printf("Length of A1: %v\n", len(a1))
	fmt.Printf("Capacity of A1: %v\n", cap(a1))
	a1 = append(a1, 1)
	fmt.Println(a1)
	fmt.Printf("Length of A1: %v\n", len(a1))
	fmt.Printf("Capacity of A1: %v\n", cap(a1))



	b1:= []int{1,2,3,4,5}
	fmt.Println(b1)
	c1 := append(b1[:2], b1[3:]...)
	fmt.Println(c1)

	//underlying array affected due to the above slice operation
	fmt.Println(b1)



}