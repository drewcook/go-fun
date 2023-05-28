// Race conditions are important to prevent and can be managed using mutex groups
// There are multiple threads/goroutines happening and only one memory space to access, and handling race conditions is essential to writing goroutines, this is solved by locking/unlocking using RWMutex and Mutex
// run `go run --race main.go` to view the race conditions
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions in golang")

	// Define our wait group
	wg := &sync.WaitGroup{}
	// Define our mutex so we prevent race conditions
	mut := &sync.RWMutex{} // best used during writes, but reads are okay too just not as important

	// We will write into our scores store while also access them
	// Lock when we need to read
	var scores = []int{0}

	// Create some goroutines using IIFEs
	wg.Add(4) // add in our three routines in one call
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One added")
		mut.Lock() // lock before any write operation
		scores = append(scores, 1)
		mut.Unlock() // unlock after accessing what is needed
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Reading scores")
		mut.RLock() // only lock when trying to read it
		fmt.Println(scores)
		mut.RUnlock() // then unlock after
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two added")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three added")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// Wait for all the routines that have been added into the wait group finish executing
	wg.Wait()
	fmt.Println(scores)
}
