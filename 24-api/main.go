// We want to build an CRUD API around music albums
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Enum for genres
type Genre int64

// Set default value, iota will automap type and increment 0, 1, 2, 3
const (
	Undefined Genre = iota
	Rock
	Pop
	HipHop
	Jazz
)

// Add a string method which will return a human readable name rather than int
func (g Genre) String() string {
	switch g {
	case Rock:
		return "rock"
	case Pop:
		return "pop"
	case HipHop:
		return "hiphop"
	case Jazz:
		return "jazz"
	}
	return "unknown"
}

// Model for albums
type Album struct {
	AlbumID      string   `json:"albumid"`
	Title        string   `json:"title"`
	Artist       string   `json:"artist"`
	Genre        Genre    `json:"genre"`
	YearReleased int      `json:"year"`
	Tracks       []string `json:"tracks"`
}

// Model for users
type User struct {
	FullName string   `json:"fullname"`
	Library  *[]Album `json:"library"`
}

// Create a mock DB
var albums []Album

func SeedAlbums() {
	// Seed our mock DB with Album data
	sgtPeppers := Album{AlbumID: "1", Title: "Sgt. Pepper's Lonely Hearts Club Band", Artist: "The Beatles", Genre: Rock, YearReleased: 1967, Tracks: []string{"Sgt. Pepper's Lonely Hearts Club Band", "With A Little Help From My Friends", "Lucy In The Sky With Diamonds", "Getting Better", "Fixing A Hole", "She's Leaving Home", "Being For The Benefit Of Mr. Kite", "Within You Without You", "When I'm Sixty Four", "Lovely Rita", "Good Morning Good Morning", "Sgt. Pepper's Lonely Hearts Club Band (Reprise)", "A Day In The Life"}}
	kindOfBlue := Album{AlbumID: "2", Title: "Kind Of Blue", Artist: "Miles Davis", Genre: Jazz, YearReleased: 1959, Tracks: []string{"So What", "Freddie Freeloader", "Blue In Green", "All Blues", "Flamenco Sketches"}}
	dsom := Album{AlbumID: "3", Title: "Dark Side Of The Moon", Artist: "Pink Floyd", Genre: Rock, YearReleased: 1973, Tracks: []string{"Speak To Me", "Breathe", "On The Run", "Time", "Great Gig In The Sky", "Money", "Us And Them", "Any Colour You Like", "Brain Damage", "Eclipse"}}
	thriller := Album{AlbumID: "4", Title: "Thriller", Artist: "Michael Jackson", Genre: Pop, YearReleased: 1982, Tracks: []string{"Wanna Be Startin' Somethin'", "Baby Of Mine", "The Girl Is Mine", "Thriller", "Beat It", "Billie Jean", "Human Nature", "P.Y.T. (Pretty Young Thing)", "The Lady In My Life"}}
	stankonia := Album{AlbumID: "5", Title: "Stankonia", Artist: "Outkast", Genre: HipHop, YearReleased: 2000, Tracks: []string{"Gasoline Dreams", "So Fresh, So Clean", "Ms. Jackson", "Snappin' & Trappin'", "Spaghetti Junction", "I'll Call Before I Come", "B.O.B", "Xplosion", "We Luv Deez Hoez", "Humble Mumble", "Drinkin' Again", "Red Velvet", "Gangsta Shit", "Toilet Tisha", "Slum Beautiful", "Stankonia (Stank Love)"}}
	albums = append(albums, sgtPeppers, kindOfBlue, dsom, thriller, stankonia)
}

// Middleware - check that an album should have an ID and title
func (a *Album) IsEmpty() bool {
	// return a.AlbumID != "" && a.Title != ""
	return a.Title == ""
}

func main() {
	fmt.Println("Building an API in golang")
	// Seed
	SeedAlbums()
	// Create a router
	router := mux.NewRouter()
	// handlers
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/albums", getAllAlbums).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	router.HandleFunc("/album", createAlbum).Methods("POST")
	router.HandleFunc("/album/{id}", updateAlbum).Methods("PUT")
	router.HandleFunc("/album/{id}", deleteAlbum).Methods("DELETE")
	// Listen
	log.Fatal(http.ListenAndServe(":4000", router))
}

// Controllers - defining routes and their handlers

// Home route handler
func serveHome(w http.ResponseWriter, r *http.Request) {
	html := `
		<h1>Welcome to the Albums API</h1>
		<p>You can view albums and add them to your library</p>
	`
	w.Write([]byte(html))
}

// Get all albums
func getAllAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all albums")
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// return all albums from our mock DB, encode them into json and write them in to the ResponseWriter
	json.NewEncoder(w).Encode(albums)
}

// Get an album by ID
// loop through our slice and return which Album matches the ID passed
func getAlbumById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get album by ID")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, album := range albums {
		if album.AlbumID == params["id"] {
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	// Not found
	json.NewEncoder(w).Encode("No album found with given ID")
	return
}

// Create a new album
func createAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new album")
	w.Header().Set("Content-Type", "application/json")

	// Validate that the request body is not empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Payload is empty")
	}

	// Decode the incoming payload into our struct
	var album Album
	_ = json.NewDecoder(r.Body).Decode(&album)
	if album.IsEmpty() {
		json.NewEncoder(w).Encode("'Title' field is required")
		return
	}

	// Check that Title is unique
	params := mux.Vars(r)
	for _, album := range albums {
		if album.Title == params["title"] {
			json.NewEncoder(w).Encode("'Title' field must be unique")
			return
		}
	}

	// Generate the unique AlbumID field using random number, convert to string
	// Naive, does not check for previous ID existing
	rand.Seed(time.Now().UnixNano())
	album.AlbumID = strconv.Itoa(rand.Intn(100))

	// Add the new course to the mock DB
	albums = append(albums, album)

	// Return the new album json
	json.NewEncoder(w).Encode(album)
	return
}

// Update an album
// loop through our slice mock DB, find matching iD, remove it from slice, update it, then add it back in
func updateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating an album")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, album := range albums {
		if album.AlbumID == params["id"] {
			// we have a match, remove the index from the slice
			albums = append(albums[:idx], albums[idx+1:]...)
			// update the fields of it
			var updatedAlbum Album
			_ = json.NewDecoder(r.Body).Decode(&album)
			updatedAlbum.AlbumID = params["id"]
			// TODO: update individual fields
			// add it back into the slice
			albums = append(albums, updatedAlbum)
			// encode it back into the response
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	// No match
	w.WriteHeader(404)
	json.NewEncoder(w).Encode("No album found with given ID")
}

// Delete an album
// loop through, find matching ID, remove it from slice
func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting an album")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var removedAlbum Album
	for idx, album := range albums {
		if album.AlbumID == params["id"] {
			removedAlbum = albums[idx]
			// we have a match, remove the index from the slice
			albums = append(albums[:idx], albums[idx+1:]...)
			break // no need to loop any further
		}
	}
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(removedAlbum)
}
