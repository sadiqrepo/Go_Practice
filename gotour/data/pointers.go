package main

import "fmt"

func main(){
	i, j := 42, 270
	var p *int = &i 
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 5
	fmt.Println(j)
}