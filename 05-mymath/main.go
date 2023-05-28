package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Working with math in golang")

	// Cannot add between different number types
	// var numA int = 2
	// var numB float64 = 4
	// fmt.Println("The sum is: ", numA+numB)

	// Randomness - both math/rand and crypto/rand are good packages to use
	// using 'math/rand
	rand.Seed(time.Now().UnixNano()) // Need to seed the algorithm, ensure it is always random (can use time)
	fmt.Println(rand.Intn(5))        // range is always exclusive, 0-4
	fmt.Println(rand.Intn(5) + 1)    // range of 1-5

	// using 'crypto/rand' - doesn't need a seed, is much more accurate also
	randomInt, _ := crypto.Int(crypto.Reader, big.NewInt(5)) // range of 0-4
	fmt.Println(randomInt)
}
