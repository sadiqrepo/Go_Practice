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


	if len(args) != 1{
		fmt.Println(usage)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}

	min := 10
	if guess > min {
		min = guess
	}

	for turn := 0 ; turn < maxTurns; turn++ {

	

	
			n := rand.Intn(min+1)
			fmt.Printf("n = %d",n)
		

		

		if n == guess {
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