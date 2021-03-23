package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		//important to close the body - if you forget - it will cause a resoruce leak
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	fmt.Println(strings.Repeat("#", 20))
	wg.Done()
}

func main() {
	urls := []string{"https://www.google.com", "https://www.subscripify.com", "https://auth-dev.subscripify.com", "https://www.micorsoft.com"}
	fmt.Println("main execution started")
	fmt.Println("No of CPUs:", runtime.NumCPU())
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup

	wg.Add(len(urls))
	for _, url := range urls {
		go checkAndSaveBody(url, &wg)

	}

	fmt.Println("No of Goroutines after start:", runtime.NumGoroutine())
	wg.Wait() //this will delay the completion of the main routine until
	fmt.Println("main execution stopped")

}
