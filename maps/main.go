package main

import "fmt"

func main() {
	var employees map[string]string

	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs:%d\n", len(employees))

	fmt.Printf(("The value for key %q is %q\n"), "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	//cannot do this = can only use comparable types - for instance slices are not comparable - but arrays are
	var m1 map[[6]int]string

	_ = m1

	//employees["Will"] = "Programmer"

	//This is a declared and initialized map
	people := map[string]float64{}

	people["Will"] = 21.4
	people["Linh"] = 10
	fmt.Println(people)

	//this will initialize an empty map
	map1 := make(map[int]int)
	fmt.Println(map1)
	map1[4] = 8
	fmt.Println(map1)

	//can also do this using map literals
	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13,
		//50:34.34,
		"CHF": 3243.1, //you do not need a comma on the last entry if they are initalized using a single line
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	//
	fmt.Println(balances["RON"]) //this will return "0" because it is a non existant key - but
	//what about zero values?
	//use the comma ok idiom
	balances["RON"] = 0
	v, ok := balances["RON"]
	if ok {
		fmt.Println("The RON balance is:", v)
	} else {
		fmt.Println("The RON Key does not exist in the map!")
	}
}
