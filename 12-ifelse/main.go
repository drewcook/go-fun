package main

import "fmt"

func main() {
	fmt.Println("if/else flow in golang")

	loginCount := 22
	var result string

	if loginCount > 20 {
		result = "heavy user"
	} else if loginCount > 10 {
		result = "regular user"
	} else {
		result = "light user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	// web request handling syntax example
	// define variable and check it in the same line
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// common to check from comma/ok syntax
	// if err != nil {

	// }
}
