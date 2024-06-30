package database

import (
	"Go-sumon/structure"
	"fmt"
	"time"

	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllBid(t *testing.T) {
		// Arrange
		ClearCollection("bid")

		postTime := time.Date(2024, 02, 29, 15, 47, 59, 264000000, time.UTC)
		// Insert some test data into the "bid" collection
		bids := []structure.Bid{}
		for i := 0; i < 2; i++ {
			id := primitive.NewObjectID() // Generate a new ObjectID
			bid := structure.Bid{
				ID:          id, // Assign the ObjectID directly
				Description: fmt.Sprintf("Bid %d", i+1),
				Time:        fmt.Sprintf("%d hours", i+1),
				BidAmount:   float64((i + 1) * 100),
				PostedTime:  postTime,
			}
			bids = append(bids, bid)

			if err := Create("bid", &bid); err != nil {
				t.Fatalf("Failed to insert bid document: %v", err)
			}
		}

		// Act: Call the GetAll function
		var resultBids []structure.Bid
		if err := GetAll("bid", &resultBids); err != nil {
			t.Fatalf("Failed to retrieve documents from the bid collection: %v", err)
		}

		// Assert: Check if the retrieved documents match the inserted ones
		if len(resultBids) != len(bids) {
			t.Fatalf("Expected %d documents, got %d", len(bids), len(resultBids))
		}
		for i, b := range resultBids {
			if !reflect.DeepEqual(b, bids[i]) {
				t.Errorf("Unexpected bid data at index %d, got %v, expected %v", i, b, bids[i])
			}
		}
	}


func TestCreateBid(t *testing.T) {
		// Arrange
		ClearCollection("bid")

		// Define a bid document to insert
		postTime := time.Date(2024, 02, 29, 15, 47, 59, 264000000, time.UTC) // Corrected time with truncated precision
		bid := structure.Bid{
			Description: "Bid for project XYZ",
			Time:        "2 hours",
			BidAmount:   200,
			PostedTime:  postTime,
		}

		// Act: Call the Create function to insert the bid document
		if err := Create("bid", &bid); err != nil {
			t.Fatalf("Failed to insert document: %v", err)
		}

		// Assert: Retrieve the inserted document and compare with the expected values
		var insertedBid structure.Bid
		if err := Get("bid", &insertedBid, bid.ID.Hex()); err != nil {
			t.Fatalf("Failed to retrieve inserted document: %v", err)
		}

		// Check if the ID field of the inserted document is set
		if insertedBid.ID.IsZero() {
			t.Error("Expected ID field to be set in the inserted document")
		}

		// Truncate the PostedTime in the inserted bid for comparison
		insertedBid.PostedTime = insertedBid.PostedTime.Truncate(time.Millisecond)

		// Check if the inserted document matches the expected values
		expectedBid := bid
		if !reflect.DeepEqual(insertedBid, expectedBid) {
			t.Errorf("Unexpected bid data, got %v, expected %v", insertedBid, expectedBid)
		}
	}

	
func TestGetBid(t *testing.T) {
		// Arrange
		collectionName := "bid"
		ClearCollection(collectionName)
		// Insert two bid documents
		postTime := time.Date(2024, 02, 29, 15, 47, 59, 264000000, time.UTC)
		expectedBid1 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project XYZ",
			Time:        "2 hours",
			BidAmount:   200,
			PostedTime:  postTime,
		}
		expectedBid2 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project ABC",
			Time:        "3 hours",
			BidAmount:   300,
			PostedTime:  postTime,
		}
		if err := Create(collectionName, &expectedBid1); err != nil {
			t.Fatalf("Failed to insert test bid document 1: %v", err)
		}
		if err := Create(collectionName, &expectedBid2); err != nil {
			t.Fatalf("Failed to insert test bid document 2: %v", err)
		}

		// Act
		var resultBid structure.Bid
		err := Get(collectionName, &resultBid, expectedBid1.ID.Hex())

		// Assert
		if err != nil {
			t.Fatalf("Failed to retrieve bid document: %v", err)
		}
		if !reflect.DeepEqual(resultBid, expectedBid1) {
			t.Errorf("Retrieved bid document does not match expected document")
		}
	}

func TestUpdateBid(t *testing.T) {
		// Arrange
		collectionName := "bid"
		ClearCollection(collectionName)

		// Insert a bid document
		expectedBid := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project XYZ",
			Time:        "2 hours",
			BidAmount:   200,
			PostedTime:  time.Now(),
		}
		if err := Create(collectionName, &expectedBid); err != nil {
			t.Fatalf("Failed to insert test bid document: %v", err)
		}

		// Define the update data
		updateData := bson.M{
			"description": "Updated bid description",
			"time":        "3 hours",
			"timeAmount":  300,
		}

		// Act: Update the bid document
		if err := Update(collectionName, expectedBid.ID.Hex(), updateData); err != nil {
			t.Fatalf("Failed to update bid document: %v", err)
		}

		// Retrieve the updated document
		var updatedBid structure.Bid
		err := Get(collectionName, &updatedBid, expectedBid.ID.Hex())

		// Assert: Verify if the retrieved document matches the expected values
		if err != nil {
			t.Fatalf("Failed to retrieve updated bid document: %v", err)
		}

		// Output the updated document with only the updated values
		fmt.Println("_id:", updatedBid.ID.Hex())
		fmt.Println("description:", updatedBid.Description)
		fmt.Println("time:", updatedBid.Time)
		fmt.Println("bidamount:", updatedBid.BidAmount)
	}

func TestDeleteBid(t *testing.T) {
		// Arrange
		collectionName := "bid"
		ClearCollection(collectionName)
		// Insert two bid documents
		postTime := time.Date(2024, 02, 29, 15, 47, 59, 264000000, time.UTC)
		expectedBid1 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project XYZ",
			Time:        "2 hours",
			BidAmount:   200,
			PostedTime:  postTime,
		}
		expectedBid2 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project ABC",
			Time:        "3 hours",
			BidAmount:   300,
			PostedTime:  postTime,
		}
		if err := Create(collectionName, &expectedBid1); err != nil {
			t.Fatalf("Failed to insert test bid document 1: %v", err)
		}
		if err := Create(collectionName, &expectedBid2); err != nil {
			t.Fatalf("Failed to insert test bid document 2: %v", err)
		}

		// Act: Call the Delete function for the second bid document
		err := Delete(collectionName, expectedBid2.ID.Hex())

		// Assert
		if err != nil {
			t.Fatalf("Failed to delete bid document: %v", err)
		}
		// Attempt to retrieve the deleted document
		var resultBid structure.Bid
		err = Get(collectionName, &resultBid, expectedBid2.ID.Hex())

		// Assert: Check if the document is not found
		expectedErrMsg := fmt.Sprintf("document with ID %s not found in collection %s", expectedBid2.ID.Hex(), collectionName)
		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Expected error '%s', got '%v'", expectedErrMsg, err)
		}
	}

func TestFindBid(t *testing.T) {
		// Arrange
		collectionName := "bid"
		ClearCollection(collectionName)
		// Insert two bid documents
		postTime := time.Date(2024, 02, 29, 15, 47, 59, 264000000, time.UTC)
		expectedBid1 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project XYZ",
			Time:        "2 hours",
			BidAmount:   200,
			PostedTime:  postTime,
		}
		expectedBid2 := structure.Bid{
			ID:          primitive.NewObjectID(),
			Description: "Bid for project ABC",
			Time:        "3 hours",
			BidAmount:   300,
			PostedTime:  postTime,
		}
		if err := Create(collectionName, &expectedBid1); err != nil {
			t.Fatalf("Failed to insert test bid document 1: %v", err)
		}
		if err := Create(collectionName, &expectedBid2); err != nil {
			t.Fatalf("Failed to insert test bid document 2: %v", err)
		}

		// Act
		var resultBids []structure.Bid
		err := Find(collectionName, bson.M{}, &resultBids)

		// Assert
		if err != nil {
			t.Fatalf("Failed to retrieve bid documents: %v", err)
		}
		if len(resultBids) != 2 {
			t.Errorf("Expected 2 bid documents, got %d", len(resultBids))
		}
	}