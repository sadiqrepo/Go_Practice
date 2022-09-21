package main

import "fmt"

func main() {

	intArray := [5]int{10, 20, 30, 40, 50}

	// ======== Example - 1 ==========

	for i :=0 ; i < len(intArray); i++ {
		fmt.Printf("%d ",intArray[i])
	}

	fmt.Println()

	// ======== Example - 2 ==========
	for index, value := range intArray {
		fmt.Printf("%d:%d ",index, value)
	}

	fmt.Println()

	// ======== Example - 3 ==========
	for _, value := range intArray {
		fmt.Printf("%d ", value)
	}

	fmt.Println()

	// ======== Example - 4 ==========
	j:=0
	for range intArray {
		fmt.Printf("%d ", intArray[j])
		j++
	}

	fmt.Println()
}
