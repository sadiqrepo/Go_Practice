package main

import (
	"fmt"
)

type vtx struct{
	Lat, Long float64
}

var m map[string]vtx

func main(){
	m = make(map[string]vtx)
	m["Bell Labs"] = vtx{
		44.432, 
		-34.654,
	}

	m["Atari"] = vtx{
		14.432, 
		34.654,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}