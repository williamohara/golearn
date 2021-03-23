package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() { //"n" is called the receiver of the type "name" above for the "print" function
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func (n names) add(newName string) {
	n = append(n, newName)

}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds() //"Seconds is a receiver function built into go

	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Dan", "Mary"}

	friends.print()

	friends.add("George") //this will append George to the  slice copy in the add function - but it wont print George below - needs pointer?

	friends.print()
}
