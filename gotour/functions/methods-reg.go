package main

import (
	"fmt"
	"math"
)


type Vxt struct{
	X, Y float64
}

func Abs(v Vxt) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vxt{3,4}
	fmt.Println(Abs(v))
}