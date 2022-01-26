package main

import (
	"fmt"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	fmt.Printf("%v, %T", actorName, actorName)
	fmt.Printf("%v, %T", companion, companion)
	fmt.Printf("%v, %T", doctorNumber, doctorNumber)
	fmt.Printf("%v, %T", season, season)

}
