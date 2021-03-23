package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

//as long as the type uses the methods listed above it will implement the interface - basically if the shape acts like a shape - its a shape
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

func (r rectangle) area2() float64 { //notice how this function is not available in the interface but you can access it directly via the receiver method
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return (2 * r.height) + (2 * r.width)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter", r.perimeter())

// }

func print(s shape) {
	fmt.Println("Shape", s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter", s.perimeter())

}

func main() {
	c1 := circle{radius: 5.}
	r1 := rectangle{height: 3., width: 2.}

	// printCircle(c1)
	// printRectangle(r1)

	print(c1)
	print(r1)
}
