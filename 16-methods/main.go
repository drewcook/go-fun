package main

import "fmt"

// methods are functions that are applied to structs (golang's concept of classes in other languages)
func main() {
	fmt.Println("methods in golang")

	user := User{"Drew", "drew@dco.dev", true, 33}
	fmt.Println(user)

	// Call some methods
	user.GetStatus()
	user.SetEmail()
	fmt.Println(user.Email)
	user.SetUnderlyingEmail()
	fmt.Println(user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// define a method directly on the struct being passed in
func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

// Passing along a copy of the object if not using a pointer
func (u User) SetEmail() {
	// does not update underlying struct property since it's just working with a copy
	u.Email = "newEmail@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

// Passing in a pointer of the object allows for updating it directly
func (u *User) SetUnderlyingEmail() {
	// updates the underlying struct property
	u.Email = "newEmail@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
