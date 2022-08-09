package main

import(
	"fmt"
)

func main(){

	_, m := multi()
	
	fmt.Println(m)

}

func multi() (int, int){
	return 5, 4
}