package main

import (
	"fmt"
)


type dtx struct{
	x, y float64
}


func (d *dtx) Scale(f float64) {
	d.x = d.x * f
	d.y = d.y * f
}

func Scalefunc(d *dtx, f float64) {
	d.x = d.x * f
	d.y = d.y * f
}


func main(){

	v := dtx{3,4}
	v.Scale(10)
	Scalefunc(&v, 10)

	p := &dtx{3,4}
	p.Scale(10)
	Scalefunc(p, 10)

	fmt.Println(v, p)
	
}