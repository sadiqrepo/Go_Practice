package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	
maxTurns = 5
usage = `Welcome to the Luck number game.
The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your numbers is, the harder it gets

Wanna Play?
`
)

func main(){

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1  {
		fmt.Println(usage)
		return
	}

	var verbose bool
	if args[0] == "-v"{
		verbose = true
	}
	
	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}


	for turn := 0 ; turn < maxTurns; turn++ {

		n := rand.Intn(guess + 1)

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			if turn == 0 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}

		//fmt.Println("You Lost... Try again?")
	}
}