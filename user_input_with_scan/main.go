package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//this reads from the command line, scanner is a pointer to buffio.Scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	//when scanner is a pointer to bufio this will force a wait for input
	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text", text)
	fmt.Println("Input bytes", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You Entered: ", text)
		if text == "exit" {
			fmt.Println("Exiting Input Scan loop")
			break
		}

	}

}
