package main

import "fmt"

func main() {
	fmt.Println("defer statements in golang")

	fmt.Println("Hello")
	// defer statements are executed line by line but invoked after the other function statements return or body of parent function completes
	defer fmt.Println("World")

	// this second defer will be invoked before the previous one on a stack (LIFO)
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
	// this will print Hello Hola Mundo World

	// these looped defers will be called before "Mundo" and then "World"
	myDefer()
}

func myDefer() {
	// prints 4, 3, 2, 1, 0
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
