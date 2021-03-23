package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf(("Length: %d, Capacity: %d \n"), len(nums), cap(nums))
	//this will get rid of the old backing array as the old  one's capacity is full
	nums = append(nums, 1, 2)

	fmt.Printf(("Length: %d, Capacity: %d \n"), len(nums), cap(nums))
}
