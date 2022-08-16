package main

import (
	"os"
	"fmt"
)

func main(){

	fmt.Println("There are",len(os.Args)-1, "people!")
	fmt.Println("Hello great",os.Args[1],"!")
	fmt.Println("Hello great",os.Args[2],"!")
	fmt.Println("Hello great",os.Args[3],"!")
}