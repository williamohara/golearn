package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Codruța"
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("‰")
	fmt.Println(n, m)

	s2 := "I Love Golang"
	fmt.Println((s2[2:5]))

	s3 := "‰I ‰Lo‰ve Go‰lang"
	fmt.Println((s3[0:1]))

	rs := []rune(s3)
	fmt.Printf(("%T\n"), rs)
	fmt.Println(string(rs[0:3]))

	// Slicing a string is efficient because it reuses the same backing array
	// Slicing returns bytes not runes

	s1 = "abcdefghijkl"
	fmt.Println(s1[2:5]) // -> cde, bytes from 2 (included) to 5 (excluded)

	s2 = "中文维基是世界上"
	fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.

	// returning a slice of runes
	// 1st step: converting string to rune slice
	rs = []rune(s2)
	fmt.Printf("%T\n", rs) // => []int32

	// 2st step: slicing the rune slice
	fmt.Println(string(rs[0:3])) // => 中文维
}
