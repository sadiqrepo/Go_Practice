package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	City string `json:"city"`
}


func main(){

	json_string :=`
	{
		"firstname" : "Daniel",
		"lastname"	: "Cormier",
		"city"		: "Bangalore"
	}`


	emp1 := new(Employee)
	emp1.FirstName = "Tony"
	emp1.LastName = "Ferguson"
	emp1.City = "Chennai"
	jsonStr, _ := json.Marshal(emp1)
	fmt.Printf("%s\n", jsonStr)

	emp2 := new(Employee)
	json.Unmarshal([]byte(json_string))
}