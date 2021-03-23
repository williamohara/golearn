package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal("This is an error log message", err)
	}
	//this will defer closing until execution is complete
	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write((byteSlice))
	if err != nil {
		log.Fatal("This is an error log message", err)
	}
	log.Printf(("Bytes Written: %d\n"), byteWritten)

	bs2 := []byte("Go Programming is cool!")

	err = ioutil.WriteFile("c.txt", bs2, 0644)
	if err != nil {
		log.Fatal("This is an error log message", err)
	}
}
