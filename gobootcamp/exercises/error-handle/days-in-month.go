package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){

	if len(os.Args) !=2 {
		fmt.Println("Give me a month name")
		return
	}
	
	year := time.Now().Year()
	leap := year % 4 == 0 && (year % 100 != 0 || year % 100 == 0)

	month := strings.ToLower(os.Args[1])

	days := 28

	if  month == "april" ||
	month == "june" ||
	month == "september" ||
	month == "november" {
		days = 30
	} else if month == "january" ||
	month == "march" ||
	month == "may" ||
	month == "july" ||
	month == "august" ||
	month == "october" ||
	month == "december" {
		days = 31
	} else if month == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n",month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)

}



