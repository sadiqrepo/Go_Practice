package main

import (
	"fmt"
)

func sayHellowWorld(){
	fmt.Println("Hello World!!")
}

func add(x int, y int){
	total := 0
	total = x + y
	fmt.Println(total)
}

func Add(x int, y int) int {
	total := 0
	total = x + y
	return total
}


func rectangle(l int, b int) (area int, perimeter int){
	//var perimeter int 
	perimeter = 2 * (l +b)
	//fmt.Println("Perimeter: ", perimeter)
	
	area = (l * b)
	return

}

func main(){
	sayHellowWorld();
	add(20, 30)
	sum := Add(30, 40)
	fmt.Println(sum)

	rectangleArea, rectanglePerimeter := rectangle(20, 30)
	fmt.Println("Rectangle Area: ", rectangleArea)
	fmt.Println("Rectangle Perimeter: ", rectanglePerimeter)
}