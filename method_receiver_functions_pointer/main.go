package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).price = newPrice
	(*c).brand = newBrand
}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000) //This will not work, because the argument was passed by value - which means that the function was working on a copy
	fmt.Println(myCar)
	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar) //still no change - as it was passed by value

	//need to use pointer receiver
	(&myCar).changeCar2("Ford", 25000)
	fmt.Println(myCar)

	//this is also syntactically correct in go
	myCar.changeCar2("Seat", 25000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	//two ways
	yourCar.changeCar2("VW", 30000)
	fmt.Println(yourCar)

	(*yourCar).changeCar2("VW Convertable", 35000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}
