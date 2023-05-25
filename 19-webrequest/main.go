package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// We will request from an external URL
const url string = "https://dco.dev"

func main() {
	fmt.Println("Working with a simple GET web request in golang")

	// make a GET request
	response, err := http.Get(url)
	checkNilError(err)

	// *http.Response - we will always get back the exact data, not a copy
	fmt.Printf("Response is of type %T\n", response)

	// it is ALWAYS the caller's responsibility to close the connection from the response, best practice is to defer it
	defer response.Body.Close()

	// Read the contents from the response body and get back a bytes array
	bytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	// Convert the bytes into stringified content
	content := string(bytes)
	fmt.Println(content) // This is all of the HTML content from the URL
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
