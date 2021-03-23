package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John",
		salary: 4000,
		contactInfo: Contact{
			email:   "wohara@wohara.com",
			address: "123 Main St.",
			phone:   6505338694,
		},
	}

	fmt.Printf("%#v\n", john)

	paul := Employee{
		name:        "paul",
		salary:      4000,
		contactInfo: john.contactInfo,
	}
	fmt.Printf("%#v\n", paul)
	fmt.Println(paul.contactInfo.email)
}
