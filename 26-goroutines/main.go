package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// For demonstrating mutex - mutually exclusive locks which lock/unlock a block of memory for reading/writing by any particular goroutine where only one goroutine can access that block at any given time - lock when it is accessing it and unlock it for others to access it when done
var signals = []string{"test"}
var mut sync.Mutex // usually should use a pointer

// Define a waitgroup for our goroutines
var wg sync.WaitGroup // usually these are pointers, but not in our test file here

func main() {
	fmt.Println("Working with concurrency and go routines in golang")
	// use 'go' keyword to open up a thread and be responsible for running greeter within that thread
	// The execution will run, then come back to main()
	go greeter("Hello")
	greeter("Goodbye")

	websiteList := []string{
		"https://dco.dev",
		"https://go.dev",
		"https://coinmarketcap.com",
		"https://adobe.com",
		"https://ethereum.org",
		"https://google.com",
		"https://github.com",
	}
	// Need to wait for all these to report back so this main() doesn't exit before these threads complete
	// We can use a wait group and the 'sync' package to achieve this.
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // since we're creating one go routine per iteration, we want to add one to our wait group, increments our wait group number
	}

	wg.Wait() // usually always goes at the end of the file, saying main() can't finish until the goroutines in the wait group have completed

	// execute another call after our wait group finishes
	fmt.Println(signals)
}

// Simple example, running this on two threads
func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(50 * time.Millisecond) // naive hack to make it wait
		fmt.Println(s)
	}
}

// Example - mocked API request
func getStatusCode(endpoint string) {
	defer wg.Done() // ensure we signal to the wait group that this has completed, decrements our wait group number

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error occurred in endpoint")
	} else {
		// append to signals here, but lock the mutex while writing, then unlock so it frees up the memory block for other calls to access
		// This is super essential when multiple services are reading/writing to a shared database and access by multiple goroutines
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
