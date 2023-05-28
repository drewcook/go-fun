// MongoDB Helpers - data layer
package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/drewcook/golang-fun-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get all records
func GetAllRecords(collection *mongo.Collection) []primitive.M {
	// Find all using empty filter
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// Close the cursor last
	defer cursor.Close(ctx)
	// Loop through using our cursor and add to our array to return
	var movies []primitive.M
	for cursor.Next(ctx) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	// Return our movies
	return movies
}

// Insert one record
func InsertOneRecord(collection *mongo.Collection, movie models.Netflix) interface{} {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted record: ", inserted)
	return inserted.InsertedID
}

// Update one record
// NOTE - for simplicity, this helper is defined to explicitly set a movie to being watched
func UpdateOneRecord(collection *mongo.Collection, movieId string) interface{} {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{"_id", id}} // filter based on our movie id
	update := bson.D{{"$set", bson.D{{"watched", true}}}} // mark it as watched
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated record modified count: ", updated.ModifiedCount)
	return id
}

// Delete one record
func DeleteOneRecord(collection *mongo.Collection, movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted record count: ", deleted.DeletedCount)
	return deleted.DeletedCount
}

// Delete all records
func DeleteAllRecords(collection *mongo.Collection) int64 {
	// everything will be selected using {} as the filter
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted record count: ", deleted.DeletedCount)
	return deleted.DeletedCount
}