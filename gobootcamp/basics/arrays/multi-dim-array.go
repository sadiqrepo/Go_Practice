package main

import (
	"fmt"
)


func main(){


	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64
	for _ , grades := range students {
		for _ , grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Sum: %g\n", sum)
	fmt.Printf("Avg grade: %.1f\n", sum/N)
}