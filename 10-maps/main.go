package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	// define new map of string keys to string values
	// use new() for zero values, or make() for non-zero values
	languages := make(map[string]string)

	// setting keys
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"

	// retrieval O(1) time
	fmt.Println("List of all languages: ", languages) // map[go:golang js:javascript py:python]
	fmt.Println("js is short for: ", languages["js"]) // "javascript"

	// deleting keys
	delete(languages, "py")
	fmt.Println("List of all languages: ", languages) // map[go:golang js:javascript]

	// iterating through the map keys using a loop
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("The value is %v\n", value)
	}
}
