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

}else{
	var isSuccess = createFile()


	if isSuccess {
fmt.Println("FIle created successfulle")
	}
}


}



func createFile() bool {
	/*
		Ask user for the name of the file to be created 
			create the file 
				if file creation failed return false 
				else return true
	*/

	return true 
}


func openFile(){}