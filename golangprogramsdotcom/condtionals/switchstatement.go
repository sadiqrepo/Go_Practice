package main


import(
	"fmt"
	"time"
)

func main(){

	today := time.Now();

	//fmt.Println(today.Day())

	switch today.Day(){
		

	case 5:
		fmt.Println("Today is 5th. Clean your house")
		fallthrough
	case 10:
		fmt.Println("Today is 10th. Buy some veggies")
		fallthrough
	case 15:
		fmt.Println("Today is 15th. Buy some food.")
		fallthrough
	case 25, 28:
		fmt.Println("Visit your doctor")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day")
	}

}