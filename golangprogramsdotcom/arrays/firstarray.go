package main

import (
	"fmt"
)

func main() {

	grades := [3]int{45, 98, 75}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	students[0] = "Lisa"
	students[1] = "Arnold"
	students[2] = "Ahmed"

	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students : %v\n", len(students))


	var identityMatrix [3][3]int = [3][3]int{
		{1, 0, 0}, {0,1,0}, {0,0,1}}
	fmt.Println(identityMatrix)


	a := [...]int{1,2,3}
	//b := a
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	

}
