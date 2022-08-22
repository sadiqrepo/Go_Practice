package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = `Usage: [command] [string]
Available commands: lower, upper and title`
)

func main(){

	if len(os.Args) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	command := os.Args[1]
	str := os.Args[2]

	switch command {
	case "lower" :
		fmt.Println(strings.ToLower(str))
	case "upper" :
		fmt.Println(strings.ToUpper(str))
	case "title" :
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n",command)
	}

}