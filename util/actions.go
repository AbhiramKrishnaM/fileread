package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadFile() {}

func OpenFile() {}

/*
Delete file
*/
func DeleteFile(filename string) bool {
	err := os.Remove(filename)

	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}
}

func DeleteFolder(foldername string) bool {
	return true
}

/*
Create file or folder
*/

func CreateFile(file_name string) bool {
	/*
		create the file with existing file name
			ask user for the file extension
				if file creation failed return false
				else return true
	*/
	var ext string
	var foldername string

	extension := map[string]string{
		"go":   "go",
		"js":   "js",
		"exe":  "exe",
		"txt":  "txt",
		"php":  "php",
		"jsx":  "jsx",
		"py":   "py",
		"yml":  "yml",
		"json": "json",
		"c":    "c",
		"psd":  "psd",
		"cpp":  "cpp",
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

func CreateFolder() bool {
	return true
}
