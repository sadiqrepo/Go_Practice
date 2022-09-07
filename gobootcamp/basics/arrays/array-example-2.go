package main

import (
	"fmt"
)

func main(){

	volumes := [...]string{
		"Kafka's revenge",
		"Stay golden",
		"Grokking python",
		"Kafka's revenge 2nd Edition",
	}

	fmt.Printf("Volumes: %#v\n", volumes)
}