package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferWrite := bufio.NewWriter(file)

	bytesWritten, err := bufferWrite.WriteString("The Go gopher is an iconic mascot!")
	if err != nil {
		log.Fatal(err)
	}
	bufferWrite.Flush()
	fmt.Println(bytesWritten, " bytes written ")
}
