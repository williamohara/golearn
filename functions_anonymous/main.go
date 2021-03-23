package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
func b() int {
	a := increment(20)
	a()
	a()
	return a()
}
func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm a horse with no name")

	a := increment(10)

	fmt.Printf("%T\n", a)
	//a is a function that will add one to x - x is effectively a variable within main()

	a()

	fmt.Println(a())

	fmt.Println(b())

}
