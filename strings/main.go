package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b' //rune literals are enclosed in single quotes, rune is an alias for int32 and it will hold the unicode value of the charater

	fmt.Printf("Char: %c, Type: %T, Value%d\n", var1, var1, var1)
	fmt.Printf("Char: %c, Type: %T, Value%d\n", var2, var2, var2)
	str := "țară"
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 0: ", str[0])
	fmt.Println("Byte (not rune) at position 1: ", str[1])
	fmt.Println("Byte (not rune) at position 2: ", str[2])
	fmt.Println("Byte (not rune) at position 3: ", str[3])
	fmt.Println("Byte (not rune) at position 4: ", str[4])
	fmt.Println("Byte (not rune) at position 5: ", str[5])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()

	for i := 0; i < len(str); {
		//need to increment to next rune by determining its size, the next
		//line returns the decoded charater and teh byte size
		//which then is used to determine how many more bytes to move forward in the string's
		//representative slice arrray. i is teh start postion
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c ", r)
		i += size
	}
	fmt.Println()

	//another way
	for _, r := range str {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	for _, r := range str {
		fmt.Printf("%T ", r)
	}
	fmt.Println()
	for _, r := range str {
		fmt.Printf("%#v ", r)
	}
	fmt.Println()
	for _, r := range str {
		fmt.Printf("%U ", r)
	}
	fmt.Println()
}
