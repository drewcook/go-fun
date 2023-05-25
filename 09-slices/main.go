package main

import "fmt"

func main() {
	fmt.Println("all about slices")

	// define, similarly like arrays, but without a specified length and initialized
	var fruitList = []string{}
	fmt.Println("fruit list is: ", fruitList)
	fmt.Printf("fruit list is: %T\n", fruitList)
	// can add to the slice, and memory is auto-adjusted
	fruitList = append(fruitList, "mango", "pineapple", "banana")
	fmt.Println("fruit list is: ", fruitList)
	fruitList = append(fruitList[1:], "pear")
	fmt.Println("fruit list is: ", fruitList)
	fruitList = append(fruitList[2:4], "strawberry")
	fmt.Println("fruit list is: ", fruitList)

}
