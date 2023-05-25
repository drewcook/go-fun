package main

import "fmt"

func main() {
	fmt.Println("We are writing functions in golang")

	// calling other functions in order
	greeter()
	greeterTwo()

	// functions can return values, and can return multiple values
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// proAdder() supports any values
	sum := proAdder(3, 4, 5, 6, 6, 38, 52)
	fmt.Println("Pro result is: ", sum)
	sum = proAdder(300, 28, 593, 0)
	fmt.Println("Pro result is: ", sum)

	// handle when multiple values are returned using comma/ok syntax
	str, count := multiReturn("hello", "world", "how", "are", "you?")
	fmt.Printf("The concatenated string of %v words is: \"%s\"\n", count, str)
}

// Functions can not be defined within other functions, but can be called within other functions
func greeterTwo() {
	fmt.Println("Another method")
} // add () after {} to make it an anonymous function and IIFE func invoked(){}()

func greeter() {
	fmt.Println("Namaste from golang")
}

// define the return data type and the order they are returned
func adder(x, y int) int {
	return x + y
}

// variatic function - can take in any number of int values using spread operator
func proAdder(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func multiReturn(values ...string) (string, int) {
	concat := ""
	count := 0
	for _, s := range values {
		concat += s + " "
		count++
	}
	return concat, count
}
