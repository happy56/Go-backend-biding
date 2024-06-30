package database

import (
	"Go-sumon/structure"
	"fmt"

	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllReview(t *testing.T) {
		// Arrange
		ClearCollection("review")

		// Insert some test data into the "review" collection
		reviews := []structure.Review{}
		for i := 0; i < 2; i++ {
			id := primitive.NewObjectID() // Generate a new ObjectID
			review := structure.Review{
				ID:            id, // Assign the ObjectID directly
				Review:        fmt.Sprintf("Review %d", i+1),
				Timelines:     float64(i + 1),
				Quality:       float64(i + 2),
				Communication: float64(i + 3),
				Behavior:      float64(i + 4),
			}
			reviews = append(reviews, review)

			if err := Create("review", &review); err != nil {
				t.Fatalf("Failed to insert review document: %v", err)
			}
		}

		// Act: Call the GetAll function
		var resultReviews []structure.Review
		if err := GetAll("review", &resultReviews); err != nil {
			t.Fatalf("Failed to retrieve documents from the review collection: %v", err)
		}

		// Assert: Check if the retrieved documents match the inserted ones
		if len(resultReviews) != len(reviews) {
			t.Fatalf("Expected %d documents, got %d", len(reviews), len(resultReviews))
		}
		for i, r := range resultReviews {
			if !reflect.DeepEqual(r, reviews[i]) {
				t.Errorf("Unexpected review data at index %d, got %v, expected %v", i, r, reviews[i])
			}
		}
	}

func TestCreateReview(t *testing.T) {
		// Arrange
		ClearCollection("review")

		// Define a review document to insert
		review := structure.Review{
			Review:        "Great service onee",
			Timelines:     4.5,
			Quality:       4.2,
			Communication: 4.8,
			Behavior:      4.6,
		}

		// Act: Call the Create function to insert the review document
		if err := Create("review", &review); err != nil {
			t.Fatalf("Failed to insert document: %v", err)
		}

		// Assert: Retrieve the inserted document and compare with the expected values
		var insertedReview structure.Review
		if err := Get("review", &insertedReview, review.ID.Hex()); err != nil {
			t.Fatalf("Failed to retrieve inserted document: %v", err)
		}

		// Check if the ID field of the inserted document is set
		if insertedReview.ID.IsZero() {
			t.Error("Expected ID field to be set in the inserted document")
		}

		// Check if the inserted document matches the expected values
		expectedReview := review
		if !reflect.DeepEqual(insertedReview, expectedReview) {
			t.Errorf("Unexpected review data, got %v, expected %v", insertedReview, expectedReview)
		}
	}

func TestGetReview(t *testing.T) {
		// Arrange
		collectionName := "review"
		ClearCollection(collectionName)
		// Insert two review documents
		expectedReview1 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service one",
			Timelines:     4.5,
			Quality:       4.2,
			Communication: 4.8,
			Behavior:      4.6,
		}
		expectedReview2 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service two",
			Timelines:     4.3,
			Quality:       4.1,
			Communication: 4.7,
			Behavior:      4.5,
		}
		if err := Create(collectionName, &expectedReview1); err != nil {
			t.Fatalf("Failed to insert test review document 1: %v", err)
		}
		if err := Create(collectionName, &expectedReview2); err != nil {
			t.Fatalf("Failed to insert test review document 2: %v", err)
		}

		// Act
		var resultReview structure.Review
		err := Get(collectionName, &resultReview, expectedReview1.ID.Hex())

		// Assert
		if err != nil {
			t.Fatalf("Failed to retrieve review document: %v", err)
		}
		if !reflect.DeepEqual(resultReview, expectedReview1) {
			t.Errorf("Retrieved review document does not match expected document")
		}
	}

func TestUpdateReview(t *testing.T) {
		// Arrange
		collectionName := "review"
		ClearCollection(collectionName)

		// Insert a review document
		expectedReview := structure.Review{
			Review:        "Great service",
			Timelines:     4.5,
			Quality:       4.2,
			Communication: 4.8,
			Behavior:      4.6,
		}
		if err := Create(collectionName, &expectedReview); err != nil {
			t.Fatalf("Failed to insert test review document: %v", err)
		}

		// Define the update data
		updateData := bson.M{
			"review":        "Updated service",
			"timelines":     4.3,
			"quality":       4.1,
			"communication": 4.7,
			"behavior":      4.5,
		}

		// Act: Update the review document
		if err := Update(collectionName, expectedReview.ID.Hex(), updateData); err != nil {
			t.Fatalf("Failed to update review document: %v", err)
		}

		// Retrieve the updated document
		var updatedReview structure.Review
		err := Get(collectionName, &updatedReview, expectedReview.ID.Hex())

		// Assert: Verify if the retrieved document matches the expected values
		if err != nil {
			t.Fatalf("Failed to retrieve updated review document: %v", err)
		}

		// Output the updated document with only the updated values
		fmt.Println("_id:", updatedReview.ID.Hex())
		fmt.Println("review:", updatedReview.Review)
		fmt.Println("timelines:", updatedReview.Timelines)
		fmt.Println("quality:", updatedReview.Quality)
		fmt.Println("communication:", updatedReview.Communication)
		fmt.Println("behavior:", updatedReview.Behavior)
	}

func TestDeleteReview(t *testing.T) {
		// Arrange
		collectionName := "review"
		ClearCollection(collectionName)
		// Insert two review documents
		expectedReview1 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service one",
			Timelines:     4.5,
			Quality:       4.2,
			Communication: 4.8,
			Behavior:      4.6,
		}
		expectedReview2 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service two",
			Timelines:     4.3,
			Quality:       4.1,
			Communication: 4.7,
			Behavior:      4.5,
		}
		if err := Create(collectionName, &expectedReview1); err != nil {
			t.Fatalf("Failed to insert test review document 1: %v", err)
		}
		if err := Create(collectionName, &expectedReview2); err != nil {
			t.Fatalf("Failed to insert test review document 2: %v", err)
		}

		// Act: Call the Delete function for the first review document
		err := Delete(collectionName, expectedReview1.ID.Hex())

		// Assert
		if err != nil {
			t.Fatalf("Failed to delete review document: %v", err)
		}
		// Attempt to retrieve the deleted document
		var resultReview structure.Review
		err = Get(collectionName, &resultReview, expectedReview1.ID.Hex())

		// Assert: Check if the document is not found
		expectedErrMsg := fmt.Sprintf("document with ID %s not found in collection %s", expectedReview1.ID.Hex(), collectionName)
		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error '%s', got '%v'", expectedErrMsg, err)
		}
	}

func TestFindReview(t *testing.T) {
		// Arrange
		collectionName := "review"
		ClearCollection(collectionName)
		// Insert two review documents
		expectedReview1 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service one",
			Timelines:     4.5,
			Quality:       4.2,
			Communication: 4.8,
			Behavior:      4.6,
		}
		expectedReview2 := structure.Review{
			ID:            primitive.NewObjectID(),
			Review:        "Great service two",
			Timelines:     4.3,
			Quality:       4.1,
			Communication: 4.7,
			Behavior:      4.5,
		}
		if err := Create(collectionName, &expectedReview1); err != nil {
			t.Fatalf("Failed to insert test review document 1: %v", err)
		}
		if err := Create(collectionName, &expectedReview2); err != nil {
			t.Fatalf("Failed to insert test review document 2: %v", err)
		}

		// Act
		var resultReviews []structure.Review
		err := Find(collectionName, bson.M{}, &resultReviews)

		// Assert
		if err != nil {
			t.Fatalf("Failed to retrieve review documents: %v", err)
		}
		if len(resultReviews) != 2 {
			t.Errorf("Expected 2 review documents, got %d", len(resultReviews))
		}
	}