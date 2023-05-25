package main

import "fmt"

// (cannot define outsite method with := syntax)
// jwtToken := 29011

// constants
const WalletAddress string = "0x0390139013"

func main() {
	var username string = "drew"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 3.14561093029341320941093
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "dco.dev"
	fmt.Println(website)
	// website = 3 (is not allowed, implicitly a string)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	// multiple declarations
	name, age := "drew", 33
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T \n", age)

	// public constants
	fmt.Println(WalletAddress)
	fmt.Printf("Variable is of type: %T \n", WalletAddress)
}
