package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	p.price = 300 //same thing would be (*p).price - this is a convenience feature in Golang
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity := 5
	price := 300.4
	name := "Laptop"
	sold := true

	fmt.Println("Before calling function", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling function", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling function with pointers", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	changeProduct(gift)

	fmt.Println(gift)

	changeProductByPointer(&gift)

	fmt.Println((gift))

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices) //notice that the values are incremented by the changeSlice even though you did not
	//pass a reference to the memory location - thats because slices are already just pointers themselves
	//same thing with maps - see eg below

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println(myMap)
}
