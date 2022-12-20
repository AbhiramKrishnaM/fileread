package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	fileName   string
	folderName string
	ext        string
)

func ReadFile() {}

func OpenFile() {}

/*
Create file or folder
*/
func CreateFile() {
	/*
		create the file with existing file name
			ask user for the file extension
				if file creation failed return false
				else return true
	*/

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

FIRST:

	fmt.Println("Enter a name for your file.")
	fmt.Scanf("%s", &fileName)

	for {
		fmt.Println("Extension for your new file ? \t")
		fmt.Scanf("%s", &ext)

		if _, ok := extension[ext]; ok {
			break
		} else {
			fmt.Printf("File extension not recognized. Please re enter \n")
		}
	}

	_filename := fileName + "." + ext

	_path := filepath.Join(_getCurrentDirectory(), _filename)

	if _, err := os.Stat(_path); err != nil {
		log.Print(err)
		os.Exit(1)
	} else {
		log.Println("Sorry this file exists.")

		if isExit := _answer(); !isExit {
			goto FIRST
		} else {
			fmt.Println("Exiting!")
			os.Exit(0)
		}

	}

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

}

func CreateFolder() {

CHECK:

	fmt.Printf("Enter folder name. This folder will be saved in the current working directory! \n")
	fmt.Scanf("%s", &folderName)

	_path := filepath.Join(_getCurrentDirectory(), folderName)

	if _, err := os.Stat(_path); err != nil {
		log.Print(err)

		if isExit := _answer(); !isExit {
			os.Exit(0)
		} else {
			/*
				create the folder
				and exit the code
			*/

			if err := os.Mkdir(folderName, os.ModePerm); err != nil {
				log.Print(err)
				os.Exit(1)
			} else {
				fmt.Println("Folder created successfully, exiting!")
				os.Exit(0)
			}
		}

	} else {
		fmt.Print("Sorry this folder exists.\n\n")

		if isExit := _answer(); !isExit {
			fmt.Println("Exiting!")
			os.Exit(0)
		} else {
			goto CHECK
		}

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

/*
Delete file
*/
func DeleteFile() {

}

func DeleteFolder() {

	fmt.Println("Enter name of the folder.")
	fmt.Scanf("%s", &folderName)

	_path := filepath.Join(_getCurrentDirectory(), folderName)

	if _, err := os.Stat(_path); err != nil {
		log.Print(err)
		os.Exit(1)
	} else {
		if err := os.RemoveAll(_path); err != nil {
			log.Print(err)
			os.Exit(1)
		} else {
			fmt.Println("Folder deleted successfully, Exiting!")
			os.Exit(0)
		}
	}

	/*
		check if the folder exisits in the current directory
		if it exists
		  delete the folder
		  else
		  say not found and then exit
	*/

}

/*
Ask user if he wants to proceed or exit the application
*/

func _answer() bool {
	var ans string
	truthyAnswer := map[string]string{
		"yes": "yes",
		"y":   "y",
		"Yes": "Yes",
	}

	fmt.Println("Do you wish to proceed? yes/no")
	fmt.Scanf("%s", &ans)

	if _, ok := truthyAnswer[ans]; ok {
		return true

	} else {
		fmt.Println("Exiting")
		return false
	}

}

/*
get current directory
*/

func _getCurrentDirectory() string {

	currentDirectory, err := os.Getwd()

	if err != nil {
		log.Print("Error fetching current directory. Exiting!")
		os.Exit(1)
	}

	return currentDirectory
}
