package main

import (
	"fmt"
)

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.1415
	const secondsInHour = 3600

	duration := 234 //in hours

	fmt.Printf("Duration in sectons: %v\n", duration*secondsInHour)

	x, y := 5, 0
	fmt.Println((x / y))

	const a = 5
	const b = 0
	//fmt.Println(a / b) this shows an error in editor because constants belong to the compile time variables belong to the run time

	//constant rules
	//1. you can not change a constant
	//2. can not initialize a constant at runtime
	//const power = math.Pow(2,3) //can not do this math.pow is runtime const needs to be set using a value
	const power = 4         //this is ok
	const l1 = len("hello") // this is ok too because len is a "built-in" function
	//3. can not use a variable to initialize a constant, variables belong to runtime
}
