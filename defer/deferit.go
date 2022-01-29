package main

import ("fmt")

func first(){
	fmt.Println("First function")
}

func second(){
	fmt.Println("Second functions")
}

func deferDemo(){
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}


func main(){
	defer second()
	first()
	deferDemo()
}