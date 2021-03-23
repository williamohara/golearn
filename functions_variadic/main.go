package main

import (
	"fmt"
	"strings"
)

//this is a variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v

	}
	return sum, product
}

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)
	return returnString
}

func main() {
	f1(1, 2, 3, 4)
	f1()

	nums := []int{1, 2}
	//append is an example of a built in variadic function
	nums = append(nums, 3, 4, 5)

	//you can pass a slice to a variadic function using the ... which will
	//prevent a new slice from being created by the function

	f1(nums...)
	f2(nums...)
	fmt.Println("Nums:", nums)

	s, p := SumAndProduct(2., 5., 10.)
	fmt.Println(s, p)
	m := personInformation(50, "William", "Aloysius", "O'Hara")
	fmt.Println(m)
}
