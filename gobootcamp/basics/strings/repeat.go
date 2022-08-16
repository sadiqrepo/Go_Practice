package main

import(
	"fmt"
	"strings"
	"os"
)

func main(){
	msg := os.Args[1]
	l := len(msg)
	m := strings.Repeat("!", l)

	s := m + msg + m

	s = strings.ToUpper(s)
	fmt.Println(s)

}