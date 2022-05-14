package main

import "fmt"

func main(){
	var s [2]string 
	s[0] = "Hello"
	s[1] = "World"
	fmt.Println(s)

	var primes = [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	a := [1]string{"Terminal"}
	fmt.Println(a)
	
}