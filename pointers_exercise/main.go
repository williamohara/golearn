package main

import "fmt"

func swap(x *float64, y *float64) {

	*y, *x = *x, *y
}

func main() {
	x := 10.10

	fmt.Printf("Address of x: %v\n", &x)

	ptr := &x

	fmt.Printf("Type of ptr: %#v\n", ptr)

	fmt.Printf("The address of the pointer: %v, and the value of x through pointer: %v\n", &ptr, *ptr)

	x2, y := 10, 2
	ptrx, ptry := &x2, &y

	z := *ptrx / *ptry

	fmt.Println(z)

	a, b := 5.5, 8.8

	swap(&a, &b)

	fmt.Println(a, "\n", b)

}
