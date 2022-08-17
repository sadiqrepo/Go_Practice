package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	char := os.Args[1]
	charLen := len(char)

	if charLen == 1 {
		if strings.IndexAny(char, "aeiou") != -1 {
			fmt.Printf("%q is a vowel.\n",char)
		} else if strings.IndexAny(char, "yw") != -1 {
			fmt.Printf("%q is sometimes a vowel, sometimes not.\n",char)
		} else {
			fmt.Printf("%q is consonant.\n",char)
		}
	} else {
		fmt.Println("Give me a letter.")
	} 

	
}