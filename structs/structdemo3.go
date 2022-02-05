package main

import (
	"fmt"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {

	e := Employee{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@golang.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   3000.00,
				TA:    1000.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    1500.00,
			},
			Salary{
				Basic: 20000.00,
				HRA:   7000.00,
				TA:    2000.00,
			},
		},
	}

	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

}
