package main

import (
	"fmt"
	"math"
)

type Vrtx struct{
	X, Y float64
}

func (v Vrtx) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main(){
	vt := Vrtx{3,4}
	fmt.Println(vt.Abs())

}