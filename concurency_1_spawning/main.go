package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) { //put pointer to waitgroup in the begining
	fmt.Println("f1 go routine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("Number - f1", i)
		time.Sleep(time.Microsecond * 8500000)
	}
	fmt.Println("f1 goroutine finished")
	wg.Done() //indicate to the wait group that this routine is done (subtracts from the number of go routines by 1)
}

func f2() {
	fmt.Println("f2 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("Number - f2", i)
		time.Sleep(time.Microsecond * 10500000)

	}
	fmt.Println("f2 finished")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No of CPUs:", runtime.NumCPU())
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup //creates a "wait group" for simultaneous routines - all will need to finish for program to finish

	wg.Add(1) //indicate the number of go routines to wait for, we only have one go routine - it can be 1000s if we want

	go f1(&wg) //pass the address of the wait group into the go routine
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())

	f2()
	wg.Wait() //this will delay the completion of the main routine until
	fmt.Println("main execution stopped")

}
