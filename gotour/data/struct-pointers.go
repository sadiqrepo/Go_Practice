package main

import "fmt"

type sp struct{
	x int
	y int
}

func main(){
	v := sp{1,2}
	p := &v
	p.x = 4

	fmt.Println(v)
}