package main

import (
	"fmt"
	"reflect"
)


func removeIndex(s []string, index int) ([]string){
	return append(s[:index], s[index+1:]...)
}

func main(){

	var intSlice []int = make([]int, 10)
	var strSlice []string = make([]string, 10, 20)

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	fmt.Println()


	fmt.Printf("intSlice \tLength: %v, \tCapacity: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLength: %v, \tCapacity: %v\n", len(strSlice), cap(strSlice))
	fmt.Println()

	// Initialize using a slice literal

    iSlice :=[]int{10, 20, 30, 40, 50}
	var sSlice []string = []string{"India", "Germany", "Singapore"}

	fmt.Printf("Items of intSlice: %v\n", iSlice)
	fmt.Printf("iSlice \tLength: %v, \tCapacity: %v\n", len(iSlice), cap(iSlice))
	fmt.Printf("Items of strSlice: %v\n", sSlice)
	fmt.Println()


	var t1slice = new([50]int)[0:10]
	fmt.Printf("t1slice \tLength: %v, \tCapacity: %v\n", len(t1slice), cap(t1slice))
	fmt.Printf("Items of t1slice: %v\n", t1slice)
	fmt.Println()

	iSlice = append(iSlice, 60, 70, 80, 90, 100)
	fmt.Printf("Items of intSlice: %v\n", iSlice)
	fmt.Printf("iSlice \tLength: %v, \tCapacity: %v\n", len(iSlice), cap(iSlice))
	fmt.Println()

	fmt.Println(iSlice[0])
	fmt.Println(iSlice[1])
	fmt.Println(iSlice[0:4])



	strSlice1 := []string{ "India", "US", "UK", "Germany", "Singapore"	}
	fmt.Printf("Items of strSlice1: %v\n", strSlice1)
	fmt.Printf("strSlice1 \tLength: %v, \tCapacity: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Println()
	
	strSlice1 = removeIndex(strSlice1, 3)
	fmt.Printf("Items of strSlice1: %v\n", strSlice1)
	fmt.Printf("strSlice1 \tLength: %v, \tCapacity: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Println()
	

}