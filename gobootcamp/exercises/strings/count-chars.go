package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)


func main() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	length := len(os.Args[1])
	fmt.Println("Bytes:",length)

	fmt.Println("Chars:",utf8.RuneCountInString(os.Args[1]))
}