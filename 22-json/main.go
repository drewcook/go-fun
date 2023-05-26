package main

import (
	"encoding/json"
	"fmt"
)

// Create a structure, don't make it public (not exported, lowercased)
type course struct {
	Name     string `json:"courseName"` // we are defining the JSON key name here
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // we are saying to ignore it from encoding and whatever is consuming it
	Tags     []string `json:"tags,omitempty"` // we are saying if the value is null, then don't show the field
}

func main() {
	fmt.Println("Working with JSON in golang")

	// Encoding different values into valid JSON
	EncodeJson()
	// Decoding JSON data into structures we can work with in Go
	DecodeJson()
}

func EncodeJson() {
	// Define our data of slice of 'course' structs
	courses := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc123", []string{"Frontend", "JavaScript", "Frameworks"}},
		{"Golang Bootcamp", 199, "Coursera", "xyz789", []string{"Backend"}},
		{"AWS Bootcamp", 149, "Udemy", "bigPW", nil},
		{"Ruby on Rails Bootcamp", 89, "Lynda", "s3cu4epw", []string{"Backend", "Frameworks"}},
	}

	// Package this data as JSON using the Marshal() method
	finalJson, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// Use a more format friendly version of Marshal(), which prettifies it adding line breaks and tabbing
	formattedJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", formattedJson)
}

func DecodeJson() {
	// Creates some test JSON data, but we should validate this JSON first (adding a trailing comma fails in the Valid() method)
	jsonDataFromWeb := []byte(`
		{
			"courseName": "ReactJS Bootcamp",
			"Price": 299,
			"website": "Udemy",
			"tags": ["Frontend", "JavaScript", "Frameworks"]
    }
	`)
	// Validate the data, put it into our structure form using Valid() method
	var testCourse course
	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("JSON was valid")
		// We want to pass in the reference here so that we're not working with a copy
		// We use Unmarshal() to convert JSON back into a format we prefer in golang (a struct)
		json.Unmarshal(jsonDataFromWeb, &testCourse)
		fmt.Printf("%#v\n", testCourse) // this is now in the form of our course struct
	} else {
		fmt.Println("JSON was not valid!")
	}

	// Some cases where you just want to add data to a key:value pair, not create a structure each time
	// Since we may not be sure what the value datatype will be, we can use the interface keyword
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// Inspect each field in the map
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v of type %T\n", k, v, v)
	}
}
