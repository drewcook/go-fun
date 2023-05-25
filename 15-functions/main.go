package main

import "fmt"

func main() {
	fmt.Println("We are writing functions in golang")

	// calling other functions in order
	greeter()
	greeterTwo()

	// functions can return values
	result := adder(3, 5)

	fmt.Println("Result is: ", result)
}

// Functions can not be defined within other functions, but can be called within other functions
func greeterTwo() {
	fmt.Println("Another method")
} // add () after {} to make it an anonymous function and IIFE func invoked(){}()

func greeter() {
	fmt.Println("Namaste from golang")
}

// define the return data type
func adder(x, y int) int {
	return x + y
}
