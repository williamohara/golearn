package main

import "fmt"

func main() {
	//	title1, author2, year1 := "The Divine Comedy", "Dante", 1320

	type book struct {
		title  string
		author string
		year   int
	}
	//not the best way to fill in  a struct - order matters! Better to use the field names
	myBook := book{"The Divine Comedy", "Dante", 1320}

	fmt.Println(myBook)

	betterBook := book{title: "Animal Farm", author: "Orwell", year: 1945}

	fmt.Println(betterBook)

	lastBook := book{title: "Caste"}
	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Wilkerson"
	lastBook.year = 2020

	fmt.Printf("%+v\n", lastBook)

	aBook := book{title: "Caste", author: "Wilkerson", year: 2020}

	fmt.Println(aBook == lastBook)

	//this creates a brand new struct - not a reference
	nextBook := aBook
	_ = nextBook

	///this is an eg of an unnamed-type struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%v\n", diana)
	fmt.Printf("Diana's age: %v\n", diana.age)

}
