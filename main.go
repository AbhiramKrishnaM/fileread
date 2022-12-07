package main

import (
	"fmt"
	"os"
)

func main() {

	var filename string

	for _, arg := range os.Args[1:] {
		filename = arg
	}

	fmt.Println(filename)
	/*
		Check if the file exists
	*/

}
