package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	sf := strings.Fields(s)

	wc := make(map[string]int)

	for _ , val := range sf {
		fmt.Println("Keyword: ", val, "Count: ", wc[val])
		wc[val] += 1
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}