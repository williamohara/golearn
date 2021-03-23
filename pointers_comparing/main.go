package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1

	fmt.Printf("Value of p1: %v, Value in the address pointed to by v1: %v\n", p1, *p1)
	fmt.Printf("p1 address is: %v and that address is also stored in pp1, here is the value in pp1 (same as &p1): %v\n", &p1, pp1)
	fmt.Printf("the *pp1 looks like this: %v\n", *pp1)
	fmt.Printf("the **pp1 looks like this(because it is a pointer to a pointer): %v\n", **pp1)

	var p2 *int

	fmt.Printf("%#v\n", p2)
	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3)   //false
	fmt.Println(*p2 == *p3) //true

}
