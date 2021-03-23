package main

import (
	"fmt"
	"log"
	"strconv"
)

func cube(n float64) float64 {
	return n * n * n
}

func f1(n uint) (uint, uint) {

	var fp uint = 1
	var sp uint = 0
	for i := n; i > 0; i-- {
		fp = fp * i
		sp = sp + i
	}
	return fp, sp
}

func myFunc(ns string) int {

	nss := []string{ns, ns + ns, ns + ns + ns}
	intr := 0

	for _, v := range nss {
		valint, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err, ":Please enter a number as a string for myFunc")
		}
		intr = intr + valint
	}

	return intr
}

func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s = s + v
	}
	return s
}
func sum2(a ...int) (s int) {
	s = 0
	for _, v := range a {
		s = s + v
	}
	return
}

func searchItem(a string, b []string) bool {
	found := false

	for _, v := range b {
		if v == a {
			found = true
		}
	}
	return found

}

func main() {
	c := cube(10)
	fmt.Println(c)
	f, s := f1(4)
	fmt.Println(f, s)
	myint := myFunc("1")
	fmt.Println(myint)
	mysum := sum(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 12, 1, 1)
	fmt.Println(mysum)

	mysum2 := sum2(1, 2, 3, 4, 5, 6)
	fmt.Println(mysum2)

	findin := []string{"William", "Aloysius", "O'Hara"}
	findstr := "Robert"
	if searchItem(findstr, findin) {
		fmt.Printf("Found %q in %s", findstr, findin)
	} else {
		fmt.Printf("%q not found in %s", findstr, findin)
	}
}
