package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	// Reads as a string
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// error handling, convert to a number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // convert string to a float, trim the \n from he input
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
