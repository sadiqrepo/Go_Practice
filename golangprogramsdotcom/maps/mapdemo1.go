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

		fmt.Println(employeeList["Sudarshan"])

	return employeeList
	
}

var (
	groceryList map[string]int = map[string]int{"Apples":10, "Oranges": 20, "SweetLime":30}
)

func addElements(){
	
	groceryList["JackFruit"] = 40
	groceryList["Banana"] = 50
	fmt.Println(groceryList)

}

func updateElements(){
	
	groceryList["JackFruit"] = 60
	groceryList["Banana"] = 70
	fmt.Println(groceryList)
	
}


func deleteElements(){

	delete(groceryList, "Banana")
	fmt.Println(groceryList)

}

func iterateMap(){

	for key, value := range groceryList{
		fmt.Println("Key: ",key, "=> Value: ",value)
	}
}

func truncateMap(){

	goMaps1 := map[string]int{"Apples" : 5, "Oranges" : 10, "Kiwis" : 20, "Rasberries" : 25}
	goMaps2 := map[string]int{"Apples" : 5, "Oranges" : 10, "Kiwis" : 20, "Rasberries" : 25}

	fmt.Println("Before Truncation")
	fmt.Println(goMaps1)
	fmt.Println(goMaps2)

	for k := range goMaps1{
		delete(goMaps1, k)
	}

	goMaps2 = make(map[string]int)

	fmt.Println("After Truncation")
	fmt.Println(goMaps1)
	fmt.Println(goMaps2)

}


func main(){
	simpleMap()
	findMapType()
	fmt.Println(mapUsingMake())
	fmt.Println("Length of map: ",len(mapUsingMake()))
	addElements()
	updateElements()
	iterateMap()
	deleteElements()
	truncateMap()

	
}