package main

import "fmt"

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is foobar")
}

func main() {
	foo()
	bar()

	defer foo() //with defer it will execute just before ending the main function
	bar()

	defer foobar() // more than one defer will execute the functions
	// from last to first in order of call LIFO
}
