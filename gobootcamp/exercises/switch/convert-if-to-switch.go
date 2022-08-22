package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user3, user4 = "jack", "inanc"
	pass3, pass4 = "1888", "1879"
)

func main(){

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	us, pd := args[1], args[2]

	switch {
	case us != user3 && us != user4:
		fmt.Printf(errUser, us)
	case us == user3 && pd == pass3 :
		fmt.Printf(accessOK, us)
	case us == user4 && pd == pass4 :
		fmt.Printf(accessOK, us)
	default:
		fmt.Printf(errPwd, us)

	}
}