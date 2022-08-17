package main

import (
	"fmt"
	"os"
)

const (
	usage1 = "Usage: [username] [password]"
	errUser1 = "Invalid User for %q.\n"
	errPass = "Invalid Password for %q.\n"
	permOK = "Access Granted for %q.\n"
	user1, user2 = "daniel", "cormier"
	pass1, pass2 = "1988", "1989"

)

func main(){

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage1)
		return
	}

	usr, pwd := args[1], args[2]

	if usr != user1 && usr != user2{
		fmt.Printf(errUser1, usr)
	} else if usr == user1 && pwd == pass1 {
		fmt.Printf(permOK, usr)
	} else if usr == user2 && pwd == pass2 {
		fmt.Printf(permOK, usr)
	} else {
		fmt.Printf(errPass, usr)
	}

}