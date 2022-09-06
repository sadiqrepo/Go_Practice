package main

import "fmt"

func main(){
	var books [4]string 
	//fmt.Printf("books : %T\n", books)
	//fmt.Println("books : ", books)
	//fmt.Printf("books : %q\n", books)
	

	books[0] = "Kafka's revenge"
	books[1] = "Java's dominance"
	books[2] = "Python's reign"
	books[3] = books[0] + " 2nd Edition"

	fmt.Printf("books : %#v\n", books)

}