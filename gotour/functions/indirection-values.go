package main


import( 
	"fmt"
	"math"
)

type Ctx struct{
	a, b float64 
}

func (c Ctx) abs() float64 {
	return math.Sqrt(c.a*c.a + c.b*c.b)
}

func absFunc(c Ctx) float64{
	return math.Sqrt(c.a*c.a + c.b*c.b)
}


func main(){
	v := Ctx{3, 4}
	fmt.Println(v.abs())
	fmt.Println(absFunc(v))

	p := &Ctx{3,4}
	fmt.Println(p.abs())
	fmt.Println(absFunc(*p))

}