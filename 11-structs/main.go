package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// there is no inheritance in golang, no super or parent

	// create a new struct instance
	drew := User{"Drew", "drew@dco.dev", true, 33}
	fmt.Println(drew)                             // {Drew drew@dco.dev true 33}
	fmt.Printf("drew details are: %v\n", drew)    // {Drew drew@dco.dev true 33}
	fmt.Printf("drew details + are: %+v\n", drew) // {Name:Drew Email:drew@dco.dev Status:true Age:33}

	// accessing values
	fmt.Printf("Name is %v and email is %v", drew.Name, drew.Email)
}

// define Structs
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
