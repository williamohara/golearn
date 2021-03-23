package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s1 := []string{"Marty", "Moe", "Jak"}

	for _, v := range s1 {
		fmt.Println(v)
	}

	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, 6)

	a := 10
	mySlice[0] = float64(a)

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

	nums := []float64{34.43, 64.78, 103879.38726737829834752}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{1.2, 2.1}

	n = append(n, nums...)

	fmt.Println(n)

	fmt.Println("No if items inside args", len(os.Args))

	if len(os.Args) >= 2 && len(os.Args) <= 10 {
		sum := 0.
		prod := 1.
		var outerErr string

	outer:
		for _, v := range os.Args[1:] {
			if i, err := strconv.ParseFloat(v, 64); err == nil {
				sum = sum + i
				prod = prod * i
			} else {
				fmt.Println("Cannot process input, product and sum not calculated")
				sum = 0.
				prod = 0.
				outerErr = err.Error()
				break outer
			}
		}
		if !(len(outerErr) > 0) {
			fmt.Printf("Sum %v, Product %v \n", sum, prod)
		}

	} else {
		fmt.Println("You need to have between 2 and 10 numbers, you have ", (len(os.Args) - 1))
	}

	num2Sum := 0
	nums2 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	nums3 := nums2[2:(len(nums2) - 2)]
	for _, v := range nums3 {
		num2Sum = num2Sum + v
	}
	fmt.Printf("The sum of %v is %v \n", nums3, num2Sum)

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	slicedYears := append(years[:3], years[len(years)-3:]...)
	fmt.Println(slicedYears)
}
