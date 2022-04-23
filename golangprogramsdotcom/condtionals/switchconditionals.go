package main

import(
	"fmt"
	"time"
)

func main(){

	switch today := time.Now(); {

	case today.Day() < 5:
		fmt.Println("Clean your house")
	case today.Day() <=10:
		fmt.Println("Buy some veggies")
	case today.Day() == 18:
		fmt.Println("Visit your doctor")
	case today.Day() >= 25:
		fmt.Println("Buy some food")
	default:
		fmt.Println("No information available for the day.")
	
	}
}