package question

import (
	"fmt"
	"unicode/utf8"
)

func question() {

	var1, var2 := 'a', 'ă'

	fmt.Printf("For var1 - Char: %c, Type: %T, Value: %d\n", var1, var1, var1)
	fmt.Printf("For var2 - Char: %c, Type: %T, Value: %d\nSo far so good, rune is an alias for int32 and the value is the Unicode Decimal Value\n\n", var2, var2, var2)

	str := "aă"

	fmt.Printf("%v is %v bytes \nI understand this, the a takes up one byte the ă takes up two bytes\n\n", str, len(str))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("For character #%v in \"str\" (%q) Char: %c, Type: %T, Value: %d\n", i+1, str, r, r, r)
		i += size
	}
	fmt.Printf("Same as above, done differently - the loop loops through the characters in the \nstring \"str\" by determining how much to jump in the underlying slice for the string\n")
	fmt.Printf("the first iteration only goes over one position and then the next iteration \ngoes over two\n")

	fmt.Println("\nNow lets go Byte by Byte ...")
	fmt.Println("Byte (not rune) at position 0: ", str[0])
	fmt.Println("Byte (not rune) at position 1: ", str[1])
	fmt.Println("Byte (not rune) at position 2: ", str[2])
	fmt.Println("Ok, I am a little confused. Position 0 holds the unicode decimal value of \"a\"")
	fmt.Printf("but what is %v and %v to  \"ă\" ?\n", str[1], str[2])

}
