package main

import (
	"fmt"
	"strings"
)

func main(){

	

	
	
	names := [...]string{
				"Einstein",
				 "Tesla",
				 "Shepard",
				}

	distances := [...]int{ 50, 40, 75, 30, 125}

	data := [...]uint8{ 72, 69, 76, 71, 75}
	

	

	ratios := [...]float64 {3.14}

	alives := [...]bool{true, false, true, false}
	
	var zero []byte
	

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for  i := 0 ; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n",i, names[i])
	}

	fmt.Println()

	fmt.Print("distances", separator)

	for i := 0 ; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Println()

	fmt.Print("data",separator)

	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Println()

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++{
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Println()

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++{
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Println()

	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++{
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	fmt.Println()

	fmt.Printf(`
%s	
	FOR RANGES 	
%[1]s 	`, strings.Repeat("~",30))

	fmt.Println()
	fmt.Print("names", separator)
	for  i , v := range names{
		fmt.Printf("names[%d]: %q\n",i, v)
	}

	fmt.Println()

	fmt.Print("distances", separator)

	for i , v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Println()

	fmt.Print("data",separator)

	for i , v := range  data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Println()

	fmt.Print("\nratios", separator)
	for i , v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Println()

	fmt.Print("\nalives", separator)
	for i , v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	fmt.Println()

	fmt.Print("\nzero", separator)
	for i , v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}

	fmt.Println()


}