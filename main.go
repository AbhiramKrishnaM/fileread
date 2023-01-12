package main

import (
	"fmt"
	"log"
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
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(cwd)

}
