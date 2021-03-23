package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	//this code purposely causes a data race against n
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second * 10)
			n++
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second * 10)
			n--
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("Final value of n", n)
}
