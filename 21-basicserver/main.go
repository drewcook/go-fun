package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// The node server from ../node-server/index.js
const baseUrl = "http://localhost:8000"

func main() {
	fmt.Print("A basic HTTP server in golang\n\n")

	// Make a GET request
	PerformGetRequest(baseUrl)
	PerformGetRequest(baseUrl + "/get")

	// Make a POST request with JSON as payload
	// Create some fake JSON
	requestBody := strings.NewReader(`
		{
			"name": "Drew",
			"title": "golang guru",
			"website": "https://dco.dev"
		}
	`)
	PerformPostRequestJson(baseUrl+"/post", requestBody)

	// Make a POST request with form data as payload
	// Create fake form data
	formData := url.Values{}
	formData.Add("name", "Drew")
	formData.Add("skill", "golang")
	formData.Add("expertise", "guru")
	PerformPostRequestFormData(baseUrl+"/postform", formData)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest(url string) {
	resp, err := http.Get(url)
	checkNilError(err)
	defer resp.Body.Close() // always defer close the body

	fmt.Println("Status code of this request is: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	content, err := ioutil.ReadAll(resp.Body)
	checkNilError(err)

	// Use the 'strings' package to read the bytes more elegantly rather than string(bytes)
	// fmt.Println(string(content))
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content) // returns the byte length
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println("The response string is: ", responseString.String()) // easy to convert it using the strings.Builder
	fmt.Print("\n\n")
}

func PerformPostRequestJson(url string, payload *strings.Reader) {
	resp, err := http.Post(url, "application/json", payload)
	checkNilError(err)
	defer resp.Body.Close()

	fmt.Println("Status code of this request is: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	content, err := ioutil.ReadAll(resp.Body)
	checkNilError(err)
	fmt.Println(string(content))
	fmt.Print("\n\n")
}

func PerformPostRequestFormData(url string, payload url.Values) {
	resp, err := http.PostForm(url, payload)
	checkNilError(err)
	defer resp.Body.Close()

	fmt.Println("Status code of this request is: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	content, err := ioutil.ReadAll(resp.Body)
	checkNilError(err)
	fmt.Println(string(content))
	fmt.Print("\n\n")
}
