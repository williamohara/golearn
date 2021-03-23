package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2
	a = int(b)
	fmt.Println(a, b)
	//var x int

	//x = "5" //creates an error

	//** ZERO VALUES **//

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                         // initialized with 0
	var price float64                     // initialized with 0.0
	var name string                       // initialized with empty string -> ""
	var done bool                         // initialized with false
	fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
}
