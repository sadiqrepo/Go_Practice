package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------


func main() {

	fmt.Printf("%v\n", os.Args)
	fmt.Println("Path: ", os.Args[0])
}