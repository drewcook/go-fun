package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drewcook/golang-fun-mongodb/router"
)

func main()  {
	fmt.Println("MongoDB API in golang")
	// Pull in our router
	r := router.Router()
	fmt.Println("Booting up server...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}