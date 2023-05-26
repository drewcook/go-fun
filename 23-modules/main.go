package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working with modules in golang")
	greeter()

	// Use gorilla router
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// Run this program as a server using the gorilla router
	// Wrap in error logger in case something fails
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello mod users")
}

// use gorilla/mux package to serve up some HTML at the index route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang modules</h1>"))
}
