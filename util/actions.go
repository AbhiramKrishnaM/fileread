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

	// _filename := file_name + "." + ext

	/*
		check if the file already exisits in the current directory
		step 1: if it does not exist
					create the file
				else
					ask for a new file name
						check if the file exist
					if yes
						say folder already exist
						repeat step 1
	*/

	return true

}

func CreateFolder(folder_name string) bool {

	var foldername string
	var ans string

	truthyAnswer := map[string]string{
		"yes": "yes",
		"y":   "y",
		"Yes": "Yes",
	}

	currentDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return false
	} else {

		if len(folder_name) == 0 {
			fmt.Printf("Enter folder name. This will the folder name to which the file will be saved! \n")
			fmt.Scanf("%s", &foldername)
		}

		_path := filepath.Join(currentDirectory, folder_name)

		if _, err := os.Stat(_path); err != nil {
			log.Print(err)

			fmt.Println("Do you wish to proceed? yes/no")
			fmt.Scanf("%s", &ans)

			if _, ok := truthyAnswer[ans]; ok {
				fmt.Println("Hello")

			} else {
				fmt.Println("Exiting")
				os.Exit(0)
			}

		} else {
			fmt.Println("Hmmmmmmmmmmmmmm")
		}

		/*
			check if the folder already exisits in the current directory
				step 1:  if it does not exist
							create the folder
						else
							ask for a new folder name
								check if the folder exists
								if yes
									say folder already exist
									repeat step 1
		*/

	}

	return true
}

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
