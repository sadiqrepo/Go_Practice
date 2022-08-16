package main


import(
	"fmt"
	"unicode/utf8"
)

func main(){

	name := "İnanç"
	fmt.Println("Bytes: ",len(name))
	fmt.Println("Codepoints: ",utf8.RuneCountInString(name))
}