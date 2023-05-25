package main

import (
	"fmt"
	"io"        // used for writing into files
	"io/ioutil" // used for reading files
	"os"        // used for creating files
)

func main() {
	fmt.Println("Working with the filesystem in golang")

	// Writing content into a newly created file
	content := "This is some arbitrary content that needs to go in a file"
	filepath := "./myNewFile.txt"
	file, err := os.Create(filepath)
	checkNilError(err)

	// write the contents into the file
	len, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length of new file contents is: ", len)

	// always close the file after working with it, recommended to defer it so it can run after all other invocations
	defer file.Close()

	// read the file contents
	readFile(filepath)
}

// Reading a file
func readFile(filename string) {
	bytes, err := ioutil.ReadFile(filename) // always read as bytes
	checkNilError(err)

	// convert the bytes into a string format
	fmt.Println("Byte data in the file is: ", bytes)
	fmt.Println("Stringified data in the file is: ", string(bytes))
}

// Abstract out common error handling into a function
func checkNilError(err error) {
	if err != nil {
		panic(err) // this will shut down the execution of the program and show the error
	}
}
