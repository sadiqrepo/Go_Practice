package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}

	name := args[0]

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(len(moods))
	n := rand.Intn(len(moods[0]))

	fmt.Printf("%s feels %s\n",name, moods[m][n] )
}