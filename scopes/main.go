package main

//import is "file scope and thus need to be declared in every file if they are used in that file"

import (
	"fmt"
	f "fmt"
) //this is permitted because it simply creating a new alias for fmt

// constants outside functions are package scoped
const done = false

//functions are package scope
func main() {
	//constants or variables declared inside functions are local scope (only in the function )
	var task = "Running"
	f.Println(task)
	f.Printf("Done in main() is %v\n", done)
	f1()
	f2()
	fmt.Println("Bye")
}

func f1() {
	const done = true
	f.Printf("Done in f1() is %v\n", done)
}
func f2() {

	f.Printf("Done in f2() is %v\n", done)
}
