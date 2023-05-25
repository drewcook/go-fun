package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "capturing user input in golang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax || err ok
	input, _ := reader.ReadString('\n')
	fmt.Print("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
