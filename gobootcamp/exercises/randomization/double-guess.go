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

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 < 0 || guess2 < 0{
		fmt.Println("Pick a positive number")
		return
	}

	min := guess1
	if guess1 < guess2{
		min = guess2
	}


	for turn := 0 ; turn < maxTurns; turn++ {

		
		n := rand.Intn(min + 1)

		if n == guess1 || n == guess2 {
			if turn == 0 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}

		fmt.Println("You Lost... Try again?")
	}
}