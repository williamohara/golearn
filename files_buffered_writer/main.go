package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	//opens the file
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	//error handleing
	if err != nil {
		log.Fatal(err)
	}
	//closes the file after execution is complete
	defer file.Close()
	//creates a buffer and assigns it to the file
	bufferedWriter := bufio.NewWriter(file)
	//creates a string "ABC" and sends the bytes to the buffer in the next line
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)
	bytesAvailable := byte(bufferedWriter.Available())
	log.Printf("Bytes available %d\n", bytesAvailable)
	//sends a string to the buffer
	bytesWritten, err = bufferedWriter.WriteString("\nJust a Random string nkjavl[pbvkij aehkiufa")
	if err != nil {
		log.Fatal(err)
	}
	//checks the buffer size
	unFlushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffred %d\n", unFlushedBufferSize)

	//writes the buffer to the file
	bufferedWriter.Flush()

	//This line would clear the buffer before it is flushed
	//bufferedWriter.Reset(bufferedWriter)

}
