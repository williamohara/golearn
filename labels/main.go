package main

import "fmt"

func main() {
	people := [6]string{"Cathy", "Beth", "Jean", "Terry", "Billy", "Bobby"}
	men := [3]string{"Terry", "Billy", "Bobby"}

outer:
	// range iterates over the array and each iteration returns the index and the value
	for index, name := range people {
		// not interested in the index so i use _
		for _, men := range men {
			if men == name {
				fmt.Printf("Found a man: %q at index %d is a dude \n", name, index)
				//this will leave the loop once a man is found
				break outer
			}

		}
	}
	fmt.Println("out of loop")
outer2:
	// range iterates over the array and each iteration returns the index and the value
	for i, name := range people {
		//
		for j, men := range men {
			if men == name {
				fmt.Printf("Found a man: %q at index %d is a dude \n", name, i)
				//this will leave the loop once a man is found
				if j == len(men) {
					break outer2
				}
			}

		}
	}
}
