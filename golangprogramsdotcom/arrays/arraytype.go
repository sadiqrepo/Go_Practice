package main

import (
	"fmt"
	"reflect"
)

func main() {

	var intArray [3]int
	var strArray [3]string

	// type of variable
	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
	fmt.Println()

	strArray[0] = "India"
	strArray[1] = "Germany"
	strArray[2] = "Singapore"

	for index := 0; index < len(strArray); index++ {
		fmt.Println(strArray[index])
	}
	fmt.Println()

	intArray = [3]int{10, 20, 30}
	fmt.Println(intArray)

	ellipseArray := [...]int{40, 50, 60, 70}
	fmt.Println(ellipseArray)

	y := [5]int{1 : 80, 3 : 100}
	fmt.Println(y)

	fmt.Println()

	for index, value := range strArray{
		fmt.Println(index, "=>", value)
	}

	fmt.Println()

	for _, value := range strArray{
		fmt.Println(value)
	}

	fmt.Println()

	j := 0
	for range strArray{
		fmt.Println(strArray[j])
		j++
	}

	fmt.Println()


	strArray2 := strArray
	strArray3 := &strArray

	fmt.Printf("strArray before modification: %v\n", strArray)
	fmt.Printf("strArra2 before modification: %v\n", strArray2)
	fmt.Printf("strArray3 before modification: %v\n", *strArray3)

	fmt.Println()

	strArray[2] = "US"

	fmt.Printf("strArray after modification: %v\n", strArray)
	fmt.Printf("strArra2 after modification: %v\n", strArray2)
	fmt.Printf("strArray3 after modification: %v\n", *strArray3)

	fmt.Println()


	countries := [...]int{10, 20, 30, 40, 50, 60, 70}
	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])
	fmt.Printf(" 1:3 %v\n", countries[1:3])
	fmt.Printf(" 2: %v\n", countries[2:])
	fmt.Printf(" 2:5 %v\n", countries[2:5])
	fmt.Printf("0:3 %v\n", countries[0:3])
	fmt.Printf("Last element %v\n", countries[len(countries)-1])

	fmt.Printf("All elements %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

}
