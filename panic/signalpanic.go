package main

import("fmt")

func main(){

	var action int

	fmt.Println("Enter 1 for Student & 2 for Professional")
	fmt.Scanln(&action)
	
	switch action {
	case 1:
		fmt.Printf("I am a student")
	case 2:
		fmt.Printf("I am a professional")
	default:
		panic(fmt.Sprintf("This is %d", action))
	}

	fmt.Println()
	fmt.Println("Enter 1 for US & 2 for UK")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("US")
	case 2:
		fmt.Printf("UK")
	default:
		panic(fmt.Sprintf("This is %d", action))
	}
	fmt.Println()

}