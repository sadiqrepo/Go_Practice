package main

import ("fmt")

func main(){

	xSlice := []int{10,20,30}
	fmt.Printf("xSlice Length - %d and Capacity - %d\n", len(xSlice), cap(xSlice))

	ySlice := make([]int, 5, 10)
	fmt.Printf("ySlice Length - %d and Capacity - %d\n", len(ySlice), cap(ySlice))

	zSlice := new([10]int)[0:5]
	fmt.Printf("zSlice Length - %d and Capacity - %d\n", len(zSlice), cap(zSlice))

	fmt.Println("ySlice before copying: ",ySlice)
	copy(ySlice, xSlice)
	fmt.Printf("ySlice Length - %d and Capacity - %d\n", len(ySlice), cap(ySlice))
	fmt.Println("ySlice after copying: ",ySlice)

	ySlice = append(ySlice, 60, 70)
	fmt.Println("ySlice after appending: ",ySlice)

	ySlice[3] = 40
	ySlice[4] = 50
	fmt.Println("ySlice after modifying: ",ySlice)
	fmt.Printf("ySlice Length - %d and Capacity - %d\n", len(ySlice), cap(ySlice))
	
	 
	


	
}