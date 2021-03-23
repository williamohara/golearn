package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	newFile, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	fileInfo, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Status check on non-existant file", err)
		} else {
			log.Fatal(err)
		}
	} else {
		err := os.Rename(fileInfo.Name(), "data.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
