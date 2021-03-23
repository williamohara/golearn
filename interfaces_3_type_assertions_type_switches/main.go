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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var s shape = circle{radius: 2.5} // create an interface variable and initialize with circle struct
	fmt.Printf("%T\n", s)

	//s.volume() //this is an error because "volume" is not part o the interface - use type assertion
	s.(circle).volume() //this is type assertion and this line is an assertion action
	//- it could fail so it is good practice to confirm assertion worked, especially in large packages so do a check
	//the assertion action returns two values - the underlying value and a boolean val

	ball, isOk := s.(circle)
	if isOk {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	//this is a type switch
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has a circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has a rectangle type\n", value)
	}
}
