package util

import (
	"log"
	"os"
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
