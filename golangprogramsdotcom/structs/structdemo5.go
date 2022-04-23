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

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===========================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}

	return "---------------------------"
}

func main() {

	e := Employee{
		FirstName: "Jon",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   1000.00,
				TA:    500,
			},
			Salary{
				Basic: 17000.00,
				HRA:   1500.00,
				TA:    700,
			},
			Salary{
				Basic: 20000.00,
				HRA:   2000.00,
				TA:    1000,
			},
		},
	}
	fmt.Println(e.EmpInfo())

}
