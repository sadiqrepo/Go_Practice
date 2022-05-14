package main

import "fmt"

func main(){
	primes := []int{2,3,5,7,11,13}
	fmt.Println(primes)

	bin := []bool{true, false, true, false, true, false}
	fmt.Println(bin)

	s := []struct{
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
	}
	fmt.Println(s)
}