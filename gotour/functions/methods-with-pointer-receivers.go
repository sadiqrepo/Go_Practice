package main

import (
	"fmt"
	"math"
)

type Ftx struct {
	X, Y float64
}

func (f *Ftx) scale(f1 float64){
	f.X = f.X * f1
	f.Y = f.Y * f1
}

func (f *Ftx) abs() float64 {
	return math.Sqrt(f.X*f.X + f.Y*f.Y)
}

func main(){
	v := &Ftx{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n",v, v.abs())
}