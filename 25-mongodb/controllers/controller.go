// Route handling and talking to the repository layer
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/drewcook/golang-fun-mongodb/models"
	"github.com/drewcook/golang-fun-mongodb/repository"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "go-netflix"
const collName = "watchlist"

// Important to ref with a pointer
var collection *mongo.Collection

// Connect
func init() {
	// Create context - use a non-nil empty context
	ctx := context.TODO()

	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create connection
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	// Set our public 'collection' var to the provided MongoDB db/collection
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Collection instance is ready")
}

// Route handlers

// Get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	movies := repository.GetAllRecords(collection)
	json.NewEncoder(w).Encode(movies)
}

// Create new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	newId := repository.InsertOneRecord(collection, movie)
	fmt.Println("New movie created with id: ", newId)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newId)
}

// Mark movie as watched
func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updatedId := repository.UpdateOneRecord(collection, params["id"])
	fmt.Println("Movie marked as watched with id: ", updatedId)
	json.NewEncoder(w).Encode(updatedId)
}

// Delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteCount := repository.DeleteOneRecord(collection, params["id"])
	fmt.Printf("Deleted %v movie with id: %v\n", deleteCount, params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteCount := repository.DeleteAllRecords(collection)
	fmt.Printf("Deleted %v movies\n", deleteCount)
	json.NewEncoder(w).Encode(deleteCount)
}