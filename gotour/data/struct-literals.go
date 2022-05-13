package main

import "fmt"

type structLiteral struct {
	x int
	y int
}

var (
	v1 = structLiteral{1, 2}
	v2 = structLiteral{x:1}
	v3 = structLiteral{}
	p = &structLiteral{1,2}
)


func main(){
	fmt.Println(v1, p, v2, v3)
}