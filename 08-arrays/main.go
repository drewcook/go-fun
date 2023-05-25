package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	// define, must specify how many positions it will take up
	var fruitList [4]string
	// add at each index
	fruitList[0] = "apple"
	fruitList[1] = "pear"
	fruitList[3] = "banana"
	fmt.Println("Fruit list is: ", fruitList)      // blank space at fruitList[2]
	fmt.Println("Fruit list is: ", len(fruitList)) // 4
	fmt.Println("Fruit list is: ", fruitList[2])   // ''

	// another way to define all at once or with pre-disposed values
	var vegList = [5]string{"potato", "onion", "pepper"}
	fmt.Println("Veggie list is: ", vegList)
}
