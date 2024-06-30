package database

import (
	"context"
	"time"

	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName = "sumon"
)

// CollectionNamesArray represents an array of collection names.
var CollectionNamesArray = []string{"review", "bid", "payment", "user", "client", "serviceProvider", "job", "point", "questionAnswer"}

// Database represents the interface for database operations.
type Database interface {
	GetAll(collectionName string, result interface{}) error
	Create(collectionName string, document interface{}) error
	UserCreate(collectionName string, document interface{}) error
	ClientCreate(document interface{}) error
	SPCreate(collectionName string, document interface{}) error
	Get(collectionName string, id string, result interface{}) error
	Update(collectionName string, id string, update interface{}) error
	Delete(collectionName string, id string) error
	Find(collectionName string, filter interface{}, result interface{}) error
}

var collection *mongo.Collection

func initMongoClient(collName string) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// for _, collName := range CollectionNamesArray {
	//     collection = client.Database(dbName).Collection(collName)
	// }
	collection = client.Database(dbName).Collection(collName)
}

func ClearCollection(collectionName string) error {
	if !contains(CollectionNamesArray, collectionName) {
		return fmt.Errorf("collection %s does not exist", collectionName)
	}

	initMongoClient(collectionName)

	_, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("failed to clear collection: %v", err)
	}
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// func First[T any](items []T) T {
// 	return items[0]
// }

// func GetAll('bid', *Bids[])

func GetAll(collectionName string, result interface{}) error {

	initMongoClient(collectionName)

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("failed to find documents in collection %s: %v", collectionName, err)
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), result); err != nil {
		return fmt.Errorf("failed to decode documents in collection %s: %v", collectionName, err)
	}

	return nil
}

/*
	func validation(collectionName string, document interface{}, validation_function func){
		// error nil
		err := validate(document)
	}

/*
*/

func Create(collectionName string, document interface{}) error {
	// Initialize MongoDB client if not already initialized
	initMongoClient(collectionName)

	// Insert the document into the collection
	res, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return fmt.Errorf("failed to insert document into collection %s: %v", collectionName, err)
	}

	// Set the ObjectID of the inserted document
	v := reflect.ValueOf(document).Elem()
	objectIDField := v.FieldByName("ID")
	if objectIDField.IsValid() && objectIDField.CanSet() {
		objectIDField.Set(reflect.ValueOf(res.InsertedID))
	} else {
		return fmt.Errorf("ID field not found or not settable")
	}

	return nil
}

func Get(collectionName string, result interface{}, id string) error {
    // Initialize the MongoDB client and collection
    initMongoClient(collectionName)

    // Convert id string to primitive.ObjectID
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("failed to parse ID: %v", err)
    }

    // Define the filter to find the document by ID
    filter := bson.M{"_id": objID}

    // Perform the database query to find the document
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err = collection.FindOne(ctx, filter).Decode(result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return fmt.Errorf("document with ID %v not found in collection %s", id, collectionName)
        }
        return fmt.Errorf("failed to find document in collection %s: %v", collectionName, err)
    }

    return nil
}

func Update(collectionName string, id string, updateData bson.M) error {
	// Convert id string to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to parse ID: %v", err)
	}

	// Define the update operation
	update := bson.M{"$set": updateData}

	// Perform the update operation
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("failed to update document in collection %s: %v", collectionName, err)
	}

	// Check if the document was found and updated
	if result.ModifiedCount == 0 {
		return fmt.Errorf("document with ID %s not found in collection %s", id, collectionName)
	}

	return nil
}

func Delete(collectionName string, id string) error {
	// Convert id string to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to parse ID: %v", err)
	}

	filter := bson.M{"_id": objID}

	// Perform the deletion operation
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete document from collection %s: %v", collectionName, err)
	}

	// Check if the document was found and deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("document with ID %s not found in collection %s", id, collectionName)
	}

	return nil
}

func Find(collectionName string, filter interface{}, result interface{}) error {
	initMongoClient(collectionName)

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to find documents in collection %s: %v", collectionName, err)
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), result); err != nil {
		return fmt.Errorf("failed to decode documents in collection %s: %v", collectionName, err)
	}

	return nil
}
