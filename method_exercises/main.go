package main

import "fmt"

// includes the print() mehtod
type money float64

// prints out amount in #.## format
func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

//book struct contains price and title
type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	vat := b.price * .09
	return vat

}

func (b *book) discount() {
	b.price = b.price * .9
}

// method for book type
// the method should receive a pointer to book, not a value
// in Go everyhing is passed by value so functions work on copies!
func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	var bucks money
	bucks = 45.9438239389
	bucks.print()

	myBook := book{title: "Where the Wild Things Are", price: 25.}

	myBook.discount()

	fmt.Printf("VAT: %.2f, Price: %.2f\n", myBook.vat(), myBook.price)

	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change

}
