package main

import (
	"fmt"
)

//enumerated constants

const(
	a = iota
	b 
	c
)

const(
	a2 = iota
)

const(
	//errorSpecialist = iota
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%v\n", a2)

	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
}