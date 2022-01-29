package main
import "fmt"
func main() {
	var action int
    fmt.Println("Enter 1 for Student and 2 for Professional")
    fmt.Scanln(&action)
    /*  Use of Switch Case in Golang */   
    switch action {
        case 1:
            fmt.Printf("I am a  Student")
        case 2:
            fmt.Printf("I am a  Professional")
		default:
			panic("Incorrect input")
    }    
    defer func() {
    	if err := recover(); err != nil{
			fmt.Println("Recovered from error", err)	
		}    	
				
    }()
}