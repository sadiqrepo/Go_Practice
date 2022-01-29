package main


import(
	"fmt"
)


var(
	area = func(l int, b int) int{
		return l * b
	}
)

func main(){
	fmt.Println(area(20,30))

	func(a int, b int){
		fmt.Println(a * b)
	}(20, 30)


	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64{
			return (f - 32.0) * (5.0/9.0)
		}(100),
	)
}