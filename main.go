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

	/*
		Check if the file exists
		if file does not exist
			create the file ask for the extension
		else
			open the file read the contents
	*/

	if _, err := os.Stat(filename); err == nil {

		fmt.Println("File exists")

	} else {
		var isSuccess = createFile(filename)

		if isSuccess {
			fmt.Println("FIle created successfully")
		}
	}

}

func createFile(file_name string) bool {
	/*
		Ask user for the name of the file to be created
			create the file
				if file creation failed return false
				else return true
	*/

	return true
}

func openFile() bool {
	/*
		Open the file
			read the contents of the file
			and close the file
	*/
	return true
}
