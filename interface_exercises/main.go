package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var myVehicle vehicle = car{licenceNo: "B143T56", brand: "Volvo"}

	fmt.Printf("Car Brand: %q Lic Num: %q\n", myVehicle.Name(), myVehicle.License())

	var empty interface{}

	empty = 1
	fmt.Printf("%T\n", empty)
	empty = 45.87643456
	fmt.Printf("%T\n", empty)
	empty = []int{34, 23, 12}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 45)
	fmt.Printf("%v\n", empty)

	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)

}
