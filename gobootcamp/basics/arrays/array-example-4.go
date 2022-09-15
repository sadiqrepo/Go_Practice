package main

import "fmt"

func main(){
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	var books [4]string
	

	for i, v := range prev {
		books[i] += v + " 2nd Ed."
	}

	fmt.Printf("last year:\n%#v\n",prev)
	fmt.Printf("\nthis year:\n%#v\n",books)
}