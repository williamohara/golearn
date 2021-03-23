package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
