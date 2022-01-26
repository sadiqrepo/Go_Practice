package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)


	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

}
