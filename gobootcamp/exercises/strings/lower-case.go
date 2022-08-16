package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	name := strings.ToLower(os.Args[1])
	fmt.Println(name)
	
}