package main

import (
	"fmt"
)

type shape interface {
	surfaceArea() float64
}

type object interface {
	volume() float64
}

// geometry is embedding shape and object interfaces, when you include an interface inside another interface you include all the methods from the included interface
type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) surfaceArea() float64 {
	return 6 * c.edge * c.edge
}

func (c cube) volume() float64 {
	return c.edge * c.edge * c.edge
}

func measure(g geometry) (float64, float64) {
	a := g.surfaceArea()
	v := g.volume()
	return a, v
}

func (c cube) getColor() string {
	return c.color
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v)
}
