package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed constant
	const b = 6.7         //untyped constant

	const c float64 = a * b    //this is ok because a and b are both constants
	const str = "Hello" + "Go" //this is ok because coded strings are effectively constants themselves

	const d = 5 > 10
	fmt.Println(d)

}
