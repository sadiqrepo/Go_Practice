package main

import (
	"fmt"
)

func simpleMap(){

    goMaps := map[string]int{"Apples" : 5, "Oranges" : 10, "Kiwis" : 20, "Rasberries" : 25}
	fmt.Println(goMaps)
}

func findMapType(){

	var iMap = map[string]int{}
	fmt.Println(iMap)
	fmt.Printf("%T\n", iMap)
}

func mapUsingMake() map[string]int{

	employeeList := make(map[string]int)
		employeeList["Suresh"] = 17
		employeeList["Sudarshan"] = 13
		employeeList["Sadiq"] = 11

	return employeeList
	
}


func main(){
	simpleMap()
	findMapType()
	fmt.Println(mapUsingMake())
	fmt.Println("Length of map: ",len(mapUsingMake()))
}