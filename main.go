package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AbhiramKrishnaM/fileread/util"
)

func main() {

	/*
		check flag
		-cf :- create a file
		-cF := create a folder

		-d :- delete folder and file
		-r : rename folder and file

			Check if the file exists
			if file does not exist
				create the file with that extension
			else
				open the file read the contents
	*/

	arguments := make(map[int]string)

	for index, arg := range os.Args[1:] {
		arguments[index] = arg
	}

	flag := arguments[0]

	selectOption(flag, arguments)

	// if _, err := os.Stat(filename); err == nil {

	// 	if openFile(filename) {
	// 		fmt.Println("File read completely")
	// 	}

	// } else {

	// 	if createFile(filename) {
	// 		fmt.Println("File created successfully")
	// 	}
	// }

}

func selectOption(_flag string, _arg map[int]string) bool {

	switch _flag {
	/*
		create a file
	*/
	case "-cf":
		fmt.Println("Gonna create a file ")

	case "-cF":
		util.CreateFolder()

	case "-dF":
		util.DeleteFolder()
	}

	return false
}

func openFile(file_name string) bool {
	/*
		Open the file
			read the contents of the file
			and close the file
	*/

	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
		return false
	} else {

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return false
		}

		return true
	}

}
