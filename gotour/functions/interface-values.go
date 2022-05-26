package main

import (
	"fmt"
	"math"
)

type In interface{
	ML()
}

type TL struct{
	S string
}

type MFloat float64

func (t *TL) ML() {
	fmt.Println(t.S)
}

func (f MFloat) ML(){
	fmt.Println(f)
}

func describe(in In){
	fmt.Printf("%v, %T\n", in, in)
}

func main(){
	
	var in In 

	in = &TL{"Interface-values"}
	describe(in)
	in.ML()

	in = MFloat(math.Pi)
	describe(in)
	in.ML()
	
}