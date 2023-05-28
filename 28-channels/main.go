// Channels tie into goroutines and mutex groups
// They are used to help allow multiple goroutines to talk to each other
// Maybe one goroutine is waiting for just part of another goroutine to complete before it starts execution
// It is a way to hand off workload to other threads/goroutines while still executing and allow signalling between different goroutines within a wait group
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	// Create a channel and define that ints will be passed around between goroutines
	channel := make(chan int, 5) // buffered by 5 with second argument
	// Typical use in production apps are that flags are passed around in channels of true/false and not using a buffered channel

	// Define a wait group and goroutines
	wg := &sync.WaitGroup{}

	// Deadlocks - channels can only allow values being passed in unless a routine is listening on the channel

	// Define our goroutines for our channel
	wg.Add(2)
	// <-chan param means "receive only" - best practice
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(channel)                  // creates issues if closing a channel before reading from it
		val, isChannelOpen := <-channel // Can receive the value in the channel from another goroutine
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-channel) // how we receive values from the channel, cannot access it until goroutines are set up and listening on it
		wg.Done()
	}(channel, wg)
	// chan<- param means "send only" - best practice
	go func(ch chan<- int, wg *sync.WaitGroup) {
		channel <- 5   // adding a value of 5 into the channel
		channel <- 6   // will only work if another listener is happening (line 26), but there are cases when we don't want that, so we'll use a buffered channel
		close(channel) // closing channel after passing in values
		wg.Done()
	}(channel, wg)

	wg.Wait()
}
