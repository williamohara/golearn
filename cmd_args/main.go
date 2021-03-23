package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("os.Args", os.Args)
	// fmt.Println("Path", os.Args[0])
	// fmt.Println("1st Arg", os.Args[1])
	// fmt.Println("2nd Arg", os.Args[2])

	fmt.Println("No if items inside args", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
	// you can shorten the above code by doing it in a simple form
	if i, err2 := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err2)
	}
}
