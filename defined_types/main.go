package main

import "fmt"

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	x = uint(s1) //need to convert the defined type s1 to uint - even though s1's type speed is a defined value based upon uint

	var s3 speed = speed(x) //need tp convert x to the defined type speed

	_ = s3

}
