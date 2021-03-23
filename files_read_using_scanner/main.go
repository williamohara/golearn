package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//this will read the file line by line - there is also word by word and custom split functions - the default for new scanner is line by line
	scanner := bufio.NewScanner(file)

	//This line will give you a split method based upon bufio.ScanRunes, bufio.ScanLine, et al.
	scanner.Split(bufio.ScanLines)

	byteGotten := false
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		} else {
			byteGotten = true
			fmt.Println(scanner.Text())
		}

	}
	if byteGotten == true {
		fmt.Println("[EOF]")
	} else {
		fmt.Println("[EOF} reached - File is empty")
	}

}
