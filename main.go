package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var filename string

	for _, arg := range os.Args[1:] {
		filename = arg
	}

	/*

		check flag
		-c :- create a folder and file

			Check if the file exists
			if file does not exist
				create the file with that extension
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

func createFile(file_name string) bool {

	/*
		create the file with existing file name
			ask user for the file extension
				if file creation failed return false
				else return true
	*/
	var ext string
	var foldername string

	extension := map[string]string{
		"go":  "go",
		"js":  "js",
		"exe": "exe",
	}

	for {
		fmt.Println("Extension for your new file ? \t")
		fmt.Scanf("%s", &ext)

		if _, ok := extension[ext]; ok {
			break
		} else {
			fmt.Printf("File extension not recognized. Please re enter \n")
		}
	}

	currentDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return false
	} else {
		_filename := file_name + "." + ext

		for {

			fmt.Printf("Enter folder name. This will the folder name to which the file will be saved! \n")
			fmt.Scanf("%s", &foldername)

			absolutePath := currentDirectory + "/" + foldername

			_, err := os.Create(filepath.Join(absolutePath, filepath.Base(_filename)))

			if err == nil {
				break
			} else {
				if err := os.Mkdir(foldername, os.ModePerm); err != nil {
					log.Fatal(err)
				} else {
					_, err := os.Create(filepath.Join(absolutePath, filepath.Base(_filename)))

					if err != nil {
						break
					}
				}
				break
			}
		}

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
