package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var filename string

	for _, arg := range os.Args[1:] {
		filename = arg
	}

	/*
		Check if the file exists
		if file does not exist
			create the file ask for the extension
		else
			open the file read the contents
	*/

	if _, err := os.Stat(filename); err == nil {

		if openFile(filename) {
			fmt.Println("File read completely")
		}

	} else {

		if createFile(filename) {
			fmt.Println("File created successfully")
		}
	}

}

type Extension int

const (
	Exe Extension = iota
	Txt
	Jpeg
	Png
	Svg
	Js
)

func (exe Extension) Name() string {
	switch exe {
	case Exe:
		return "exe"
	case Txt:
		return "txt"
	case Jpeg:
		return "jpeg"

	case Png:
		return "png"
	case Svg:
		return "svg"
	case Js:
		return "js"

	}

	return "error"
}

func createFile(file_name string) bool {
	/*
		create the file with existing file name
			ask user for the file extension
				if file creation failed return false
				else return true
	*/
	var ext string

	fmt.Println("Extension for your new file ? \t")
	fmt.Scanf("%s", &ext)

	/*
		check if the extension is an enum
	*/

	_filename := file_name + "." + ext

	file, err := os.Create(_filename)

	if err != nil {
		log.Fatal(err)
		return false
	} else {

		fmt.Println(&file)

		return true
	}

}

func openFile(file_name string) bool {
	/*
		Open the file
			read the contents of the file
			and close the file
	*/
	return true
}
