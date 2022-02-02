package main

import ( "fmt")


type rectangle struct{
	length float64
	breadth float64
	color string
}

func main(){
	var rect1 = rectangle{10.5, 25.10, "red"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 5, color: "blue"}
	fmt.Println(rect2)

	var rect3 = rectangle{length: 10, breadth: 20, color: "green"}
	fmt.Println(rect3)

	var rect4 = rectangle{breadth: 25.10, color: "yellow"}
	fmt.Println(rect4)

	var rect5 = rectangle{breadth: 25.10}
	fmt.Println(rect5)

	var rect6 = rectangle{}
	fmt.Println(rect6)

	rect7 := new(rectangle)
	rect7.length = 5
	rect7.breadth = 10
	rect7.color = "white"
	fmt.Println(rect7)


	var rect8 = new(rectangle)
	rect8.length = 20
	rect8.breadth = 40
	rect8.color = "black"
	fmt.Println(rect8)


	var rect9 = &rectangle{10, 20, "Pink"}
	fmt.Println(rect9)

	var rect10 = &rectangle{}
	rect10.length = 10
	rect10.color = "Purple"
	fmt.Println(rect10)

	var rect11 = &rectangle{}
	(*rect11).breadth = 20
	(*rect11).color = "Magenta"
	fmt.Println(rect11)


}