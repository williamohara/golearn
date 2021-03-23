////////////////////////////////
// For Loops
// Go Playground: https://play.golang.org/p/RiErMJCR3Z_c
/////////////////////////////////

package main

import "fmt"

func main() {

	// printing numbers from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum) //this line is never reached

	/////////////////////////////////
	// For, Continue and Break Statements
	// Go Playground: https://play.golang.org/p/HUdSS9_TSLP
	/////////////////////////////////

	//** CONTINUE STATEMENT **//

	// It works just the same as in C,  Java or Python.
	// The continue statement rejects all the remaining statements in the current iteration of the loop
	// and moves the control back to the top of the loop.

	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	// **BREAK STATEMENT **//

	// It is used to terminate the innermost for or switch statement.
	// It works just the same as in C,  Java or Python.

	// finding 10 numbers divisible by 13
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	// the break statement is not terminating the program entirely;
	fmt.Println("Just a message after the for loop")
}
