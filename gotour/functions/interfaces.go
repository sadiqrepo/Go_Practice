package main

import (
	"fmt"
	"math"
)

type Abser interface {
	 Abs() float64
}

type Myfloat float64

type Vertex struct {
	X, Y float64
}

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	var a,b,c Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3,4}
	v1 := &Vertex{3,4}

	a = f
	b = &v
	c = v1
	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
	fmt.Println(c.Abs())

}