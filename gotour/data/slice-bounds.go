package main

import "fmt"

func main(){
	s := []int{2,3,5,7,11,13}

	s1 := s[0:6]
	s2 := s[:3]
	s3 := s[2:]
	s4 := s[:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}