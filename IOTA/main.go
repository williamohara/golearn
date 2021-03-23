package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)

	fmt.Println(c11, c22, c33)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Println(North, East, South, West)

	const (
		a = -(iota) - 1
		_ //skips -2 because using the _ "blank identifier"
		b
		c
	)
	fmt.Println(a, b, c)
}
