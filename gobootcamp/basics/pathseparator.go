package main

import (
	"fmt"
	"path"
)

func main(){
	//var dir, file string
	//dir, file = path.Split("css/main.css")
	//fmt.Println("dir: ", dir)

	_, file := path.Split("css/main.css")
	fmt.Println("file: ", file)
}