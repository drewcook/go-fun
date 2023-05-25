package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://dco.dev:3000/skills?skillname=golang&expertise=guru"

func main() {
	fmt.Println("Working with URLs in golang")
	fmt.Println(myurl)

	// Parsing the parts of the URL
	result, _ := url.Parse(myurl)
	fmt.Println("URL scheme is: ", result.Scheme)
	fmt.Println("URL host is: ", result.Host)
	fmt.Println("URL port is: ", result.Port())
	fmt.Println("URL path is: ", result.Path)
	fmt.Println("URL query is: ", result.RawQuery)

	// Parsing query params
	qparams := result.Query() // this puts the parameters into a map
	fmt.Println(qparams)
	fmt.Printf("The datatype of the query params is: %T\n", qparams)
	fmt.Println(qparams["skillname"])
	fmt.Println(qparams.Get("expertise")) // a more simple way to get the actual value, not [value]

	for _, val := range qparams {
		fmt.Println("Looping over, the param is: ", val)
	}

	// Constructing a URL out of different parts
	// Always work with a REFERENCE of the url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "dco.dev",
		Path:     "/about",
		RawQuery: "user=drew",
	}
	constructedUrl := partsOfUrl.String()
	fmt.Println(constructedUrl)
}
