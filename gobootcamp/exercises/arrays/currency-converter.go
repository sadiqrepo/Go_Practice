package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){

	const (
		EUR = iota
		GBP
		JPY
		INR
	)

	rates := [...]float64{
		EUR: 1.00,
		GBP: 0.88,
		JPY: 143.24,
		INR: 79.73,

	}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number!")
		return
	}

	fmt.Printf("%g USD is %.2f EUR\n", amount, amount * rates[EUR])
	fmt.Printf("%g USD is %.2f GBP\n", amount, amount * rates[GBP])
	fmt.Printf("%g USD is %.2f JPY\n", amount, amount * rates[JPY])
	fmt.Printf("%g USD is %.2f INR\n", amount, amount * rates[INR])






}