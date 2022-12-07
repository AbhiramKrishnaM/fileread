package main

import (
	"fmt"
	"os"
)

func main() {

	// var filename string

	// for _, arg := range os.Args[1:] {
	// 	filename = arg
	// }

	/*
		Check if the file exists
	*/

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(cwd)

}
