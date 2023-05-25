package main

import "fmt"

func main() {
	fmt.Println("Welcome to the lesson on pointers")

	// creating variables, which have underlying pointers, but sometimes copies get passed along through functions, classes, etc and not the actual value in memory
	var one int = 1
	fmt.Println("Value of pointer is ", one)

	// define a datatype of pointer
	var ptr *int
	fmt.Println("Value of pointer is ", ptr) // default value is nil

	// create a pointer that points to another memory addres of something that has already been declared
	myNumber := 20
	var pntr = &myNumber
	fmt.Println("Value of myNumber is ", myNumber) // is 20
	fmt.Println("Value of pointer is ", pntr)      // is 0xc00001a0b8, the actual memory address location
	fmt.Println("Value of pointer is ", *pntr)     // is 20, saying what is inside that pointer

	// redefine an existing pointer
	*pntr = *pntr * 2
	fmt.Println("New value of pointer is ", myNumber) // we are multiplying the underlying value in memory, so it is guaranteed that we are updating the actual value, which is 40
}
