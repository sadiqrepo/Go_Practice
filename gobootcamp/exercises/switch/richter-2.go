package main 

import (
	"fmt"
	"os"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var scale string 
	desc := os.Args[1]

	switch desc {

	case "massive" :
		scale = "10+"
	case "great" :
		scale = "8 - 9.9"
	case "major" :
		scale = "7 - 7.9"
	case "strong" :
		scale = "6 - 6.9"
	case "moderate" :
		scale = "5 - 5.9"
	case "light" :
		scale = "4 - 4.9"
	case "minor" :
		scale = "3 - 3.9"
	case "very minor" :
		scale = "2 - 2.9"
	case "micro" :
		scale = "less than 2.0"
	default:
		scale = "unknown"
	}

	fmt.Printf("%s's richter scale is %s\n", desc, scale)

}