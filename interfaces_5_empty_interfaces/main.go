package main

import "fmt"

type emptyInterface interface {
	// any go type satisfies this interface which means that it can represent any values. It can hold values of any type
	// fmt.Println is an example of an empty interface built into go - it can take any value of any time in any amount of values
}

type person struct {
	info interface{} //example of a struct with an empty interface
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)
	//fmt.Println(len(empty)) //this is a compile error, an interface stores the dynamic type and a dynamic concrete value - you must do an assertion
	fmt.Println(len(empty.([]int))) //assertion done this way - by asserting that it is of []int, you can get the length

	//empty interfaces are used best for situations where you do not know the type comming in

	you := person{} //as this is an empty interface - you can put anything you want into person
	you.info = "William"
	fmt.Println(you.info)
	you.info = 40
	fmt.Println(you.info)
	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info)
}
