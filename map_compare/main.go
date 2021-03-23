package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}
	//to compare maps this will not work - one can only compare a map to nil

	//fmt.Println(a == b)

	fmt.Println((fmt.Sprintf("%s", a)) == (fmt.Sprintf(("%s"), b)))
}
