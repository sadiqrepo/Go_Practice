package main

import (
	"fmt"
)

type I interface{
	M()
}

type T struct{
	S string
}

func (t T) M(){
	fmt.Println(t.S)
}

func (t T) A(){
	fmt.Println(t.S)
}

func main(){
	var i I = T{"Hello!!"}
	var j T = T{"How are you?"}

	i.M()
	j.A()
}