package main

import "fmt"

// when assigning maps they assign a reference to the original map - i.e. the same map header
func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbors := friends

	friends["Dan"] = 50

	fmt.Println(neighbors)

	//to create a real copy - initialize a new map and use a for loop

	people := make(map[string]int)

	for k, v := range neighbors {
		people[k] = v

	}
	people["Dan"] = 34
	fmt.Println(people)
	fmt.Println(neighbors)
	fmt.Println(friends)
}
