package main

import "fmt"

func change(a *int) *float64 { //takes in a pointer type and returns a pointer type to a different value type

	*a = 100 //assigns 100 to the memory location passed in

	b := 5.5 //puts a value in b

	return &b //sends the address of b back in the return value (note that a was a pointer the the value of x in the main function yet x is still chaged - no return needed)

}

func main() {
	x := 8  //assigning 8 to x
	p := &x //pointing p to x, p holds x's memory location

	fmt.Println("Value of x before calling change: ", x)
	z := change(p)
	fmt.Println("Value of x after  calling change: ", x)
	fmt.Printf("The value of z, which is a pointer to a float is %v, \nit is pointing to the memory location of \"b\" that was assigned in the \n \"change\" function which holds %v", z, *z)

}
