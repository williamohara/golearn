package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		//this is idomatic
		log.Fatal("This is an error log message", err)

	}
	//this will empty out the file all except for the firs x(2nd arg) bytes
	err = os.Truncate("a.txt", 0)

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		//this is idomatic
		log.Fatal("This is an error log message", err)

	}

	newFile.Close()

	//one way to open a file - but this will overwrite the file that is already there
	//file, err := os.Open("a.txt")
	//better way to open a file if you are not sure it exists or not
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		//this is idomatic
		log.Fatal("This is an error log message", err)

	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name:", fileInfo.Name())

	//this is an example of when an error is returned when a file does not exist

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			//log.Fatal ends the program
			//log.Fatal("File does not exist!\n", err)
			fmt.Println("File does not exist!")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"

	err = os.Rename(oldPath, newPath)
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		//this is idomatic
		log.Fatal("This is an error log message", err)

	}

	err = os.Remove("aaa.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		//this is idomatic
		log.Fatal("This is an error log message", err)

	}
}
