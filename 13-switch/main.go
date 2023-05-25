package main

import (
	"fmt"
	"math/rand"
	"time"
)

// a dice game
func main() {
	fmt.Println("switch cases in golang")

	// need a random number, defining to seed it as the current nanosecond rather than randomly on startup
	rand.Seed(int64(time.Now().UnixNano()))
	// define to a random number
	diceNumber := rand.Intn(6) + 1 // between 1 and 6
	fmt.Println("Value of dice is: ", diceNumber)

	// switch
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can open")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
		fallthrough // this says to fall through to the next switch case if this case is hit
	case 4:
		fmt.Println("You can move 4 spaces")
		fallthrough // if we roll 4, we will see log lines from case 4 and case 5
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces and roll the dice again")
	default:
		fmt.Println("What roll did you do?")
	}
}
