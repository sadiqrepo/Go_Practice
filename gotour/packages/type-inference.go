package main

import "fmt"

func main(){
	 i := 42
	 j := 42.5
	 k := 42.5 + 1i

	 fmt.Printf("I is of type %T\n", i)
	 fmt.Printf("J is of type %T\n", j)
	 fmt.Printf("K is of type %T\n", k)
}