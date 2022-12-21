package main

import (
	"os"

	"github.com/AbhiramKrishnaM/fileread/util"
)

func main() {

	/*
		check flag
		-cf :- create a file
		-cF :- create a folder

		-rf :- read a file
		-rF :- read contents of folder


		-d :- delete folder and file
		-r :- rename folder and file

	*/

	arguments := make(map[int]string)

	for index, arg := range os.Args[1:] {
		arguments[index] = arg
	}

	flag := arguments[0]

	selectOption(flag, arguments)

}

func selectOption(_flag string, _arg map[int]string) bool {

	switch _flag {
	/*
		create a file
	*/
	case "-cf":
		util.CreateFile()

	case "-cF":
		util.CreateFolder()

	case "-df":
		util.DeleteFile()

	case "-dF":
		util.DeleteFolder()

	case "-rf":
		util.ReadFile()

	case "-rF":
		util.ReadFolder()
	}

	return false
}
