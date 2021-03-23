package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}

	fmt.Println(accounts)

	names := [...]string{
		5: "William",
		"Robert",
		0: "Terrence",
		"Tom",
		"Richard",
		"Robbie",
		"Jack",
		15: "Tommy",
		32: "Andrew",
	}
	fmt.Println(names)
}
