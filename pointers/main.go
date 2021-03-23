package main

import "fmt"

func main() {
	name := "William"
	fmt.Println(&name) //this shows the memory location

	var x int = 2

	ptr := &x

	fmt.Printf("ptr is a %T with a value of %v\n", ptr, ptr)
	fmt.Printf("ptr holds the memory address of x, which is %v, and x holds a value of %v\n", &x, *ptr)
	fmt.Printf("ptr's address is stored at address: %v\n", &ptr)

	//just declairing a pointer with out initializing

	var ptr1 *float64
	_ = ptr1

	p := new(int) //created a pointer p that points to an int type

	x = 100
	p = &x //this initializes p with the value for the memory location of x

	fmt.Printf("p is a %T with a value of %v\n", p, p)
	fmt.Printf("p is a %T pointing to the memory location that holds the value of %v\n", p, *p)
}
