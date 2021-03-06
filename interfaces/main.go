package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return (2 * r.height) + (2 * r.width)
}

func printCircle(c circle) {
	fmt.Println("Shape", c)
	fmt.Println("Area:", c.area())
	fmt.Println("Perimeter", c.perimeter())
}

func printRectangle(r rectangle) {
	fmt.Println("Shape", r)
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter", r.perimeter())

}

//this is not good practice!!! you have developed two functions which do basically the same thing - you need to use interfaces
func main() {
	c1 := circle{radius: 5.}
	r1 := rectangle{height: 3., width: 2.}

	printCircle(c1)
	printRectangle(r1)
}
