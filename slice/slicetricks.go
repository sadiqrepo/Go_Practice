package main

import("fmt")

func main(){


	var countries = []string{"India", "Canada", "US", "UK", "Germany"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(" [:2] %v\n", countries[:2])
	fmt.Printf(" [1:3] %v\n", countries[1:3])
	fmt.Printf(" [2:] %v\n", countries[2:])
	fmt.Printf(" [2:5] %v\n", countries[2:5])
	fmt.Printf(" [0:3] %v\n", countries[0:3])


	fmt.Printf("Last Element: %v\n", countries[4])
	fmt.Printf("Last Element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last Element: %v\n", countries[4:])

	fmt.Printf("All Elements: %v\n", countries[0:len(countries)])
	fmt.Printf("Last 2 Elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last 2 Elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	// Looping through slices
	fmt.Println("\n ---- Example 1 ----")

	for index, value := range countries{
		fmt.Println("For index[",index,"] == ", value)
	}

	// Example 2
	fmt.Println("\n ---- Example 2 ----")
	for _ , value := range countries{
		fmt.Println(value)
	}

	// Example 3
	fmt.Println("\n ---- Example 3 ----")
	j := 0 
	for range countries{
		fmt.Println(countries[j])
		j++
	}
}