package main

import(
	"fmt"
)

func main(){
	var s string 
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	t := `
	<html>
		<body>"Hello"</body>
	</html>
	`

	fmt.Println(t)

	fmt.Println("C:\\my\\dir\\path")
	fmt.Println(`C:\\new\\dir\\path`)
}