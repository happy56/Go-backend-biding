package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"Go-sumon/database"
	"Go-sumon/structure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Insert two bid documents for testing
	bid1 := structure.Bid{
		ID:          primitive.NewObjectID(),
		Description: "Bid 1 description",
		Time:        "Bid 1 time",
		BidAmount:   100,
		PostedTime:  time.Now(),
	}
	bid2 := structure.Bid{
		ID:          primitive.NewObjectID(),
		Description: "Bid 2 description",
		Time:        "Bid 2 time",
		BidAmount:   150,
		PostedTime:  time.Now(),
	}
	if err := database.Create("bid", &bid1); err != nil {
		t.Fatalf("Failed to insert test bid document 1: %v", err)
	}
	if err := database.Create("bid", &bid2); err != nil {
		t.Fatalf("Failed to insert test bid document 2: %v", err)
	}

	// Create a new HTTP GET request
	req := httptest.NewRequest("GET", "/bid", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	GetAllBidHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	// Decode the response body into a slice of Bid objects
	var bids []structure.Bid
	if err := json.NewDecoder(rr.Body).Decode(&bids); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check if the correct number of bids was retrieved
	if len(bids) != 2 {
		t.Errorf("unexpected number of bids retrieved: got %d, want 2", len(bids))
	}

}

func TestCreateBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Create a new Bid object for the request body
	requestBody := structure.Bid{
		ID:          primitive.NewObjectID(),
		Description: "Bid 1 description",
		Time:        "Bid 1 time",
		BidAmount:   100,
		PostedTime:  time.Now(),
	}

	// Marshal the Bid object into JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	// Create a new HTTP request with the marshaled JSON as the request body
	req := httptest.NewRequest("POST", "/bid", bytes.NewReader(requestBodyJSON))

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	CreateBidHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the content type of the response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	// Decode the response body
	var createdBid structure.Bid
	if err := json.NewDecoder(rr.Body).Decode(&createdBid); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check if the bid was created successfully
	if createdBid.ID.IsZero() {
		t.Error("bid ID was not set")
	}
	if createdBid.Description != "Bid 1 description" {
		t.Errorf("unexpected bid description: got %s, want Bid 1 description", createdBid.Description)
	}
	// Check other fields as needed
}

func TestGetBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Create mock bid documents for testing
	mockBid1 := structure.Bid{
		ID:          primitive.NewObjectID(),
		Description: "Great bid 1",
		Time:        "10:00",
		BidAmount:   100,
		PostedTime:  time.Now(),
	}

	mockBid2 := structure.Bid{
		ID:          primitive.NewObjectID(),
		Description: "Great bid 2",
		Time:        "11:00",
		BidAmount:   150,
		PostedTime:  time.Now(),
	}

	// Insert the mock bid documents into the database
	if err := database.Create("bid", &mockBid1); err != nil {
		t.Fatalf("Failed to insert mock bid document 1: %v", err)
	}
	if err := database.Create("bid", &mockBid2); err != nil {
		t.Fatalf("Failed to insert mock bid document 2: %v", err)
	}

	// Create a new HTTP GET request with the ID of the second mock bid in the URL query
	req, err := http.NewRequest("GET", "/bid?id="+mockBid2.ID.Hex(), nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	GetBidHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	// Decode the response body into a Bid object
	var retrievedBid []structure.Bid
	if err := json.NewDecoder(rr.Body).Decode(&retrievedBid); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
}

func TestUpdateBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Insert a bid document for testing
	expectedBid := structure.Bid{
		Description: "Bid description", // <-- This needs to be "Updated description"
		Time:        "2024-03-14T12:00:00Z",
		BidAmount:   100,
		PostedTime:  time.Now(),
	}
	if err := database.Create("bid", &expectedBid); err != nil {
		t.Fatalf("Failed to insert test bid document: %v", err)
	}

	// Define the update data
	updateData := map[string]interface{}{
		"description": "Updated description",
		"time":        "2024-03-15T12:00:00Z",
		"bidAmount":   150,
	}

	// Marshal the update data to JSON
	updateBody, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("Failed to marshal update data: %v", err)
	}

	// Convert the ObjectID to a string
	idString := expectedBid.ID.Hex()

	// Create a PUT request with the update data
	req, err := http.NewRequest("PUT", "/bid?id="+idString, bytes.NewBuffer(updateBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Execute the handler function
	UpdateBidHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Retrieve the updated document from the database
	var updatedBid structure.Bid
	err = database.Get("bid", &updatedBid, idString)
	if err != nil {
		t.Fatalf("Failed to retrieve updated bid document: %v", err)
	}

	// Check if the bid was updated successfully
	if updatedBid.Description != "Updated description" {
		t.Errorf("unexpected bid description: got %s, want Updated description", updatedBid.Description)
	}
	// Check other fields as needed
}

func TestDeleteBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Insert a bid document for testing
	expectedBid := structure.Bid{
		Description: "Test bid",
		Time:        "2024-03-14T12:00:00Z",
		BidAmount:   100,
		PostedTime:  time.Now(),
	}
	if err := database.Create("bid", &expectedBid); err != nil {
		t.Fatalf("Failed to insert test bid document: %v", err)
	}

	// Convert the ObjectID of the document to a string
	id := expectedBid.ID.Hex()

	// Create a DELETE request with the bid ID
	req, err := http.NewRequest("DELETE", "/bid?id="+id, nil)
	if err != nil {
		t.Fatalf("Failed to create delete request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Execute the handler function
	DeleteBidHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the document still exists in the database
	var remainingBid structure.Bid
	err = database.Get("bid", &remainingBid, id)
	if err == nil {
		t.Error("Document still exists in the database after deletion")
	}
}

func TestFindBidHandler(t *testing.T) {
	// Clear the "bid" collection before running the test
	database.ClearCollection("bid")

	// Insert test bids into the database
	testBids := []structure.Bid{
		{Description: "First bid", Time: "2024-03-14T12:00:00Z", BidAmount: 100, PostedTime: time.Now()},
		{Description: "Second bid", Time: "2024-03-14T13:00:00Z", BidAmount: 150, PostedTime: time.Now()},
	}

	for _, bid := range testBids {
		if err := database.Create("bid", &bid); err != nil {
			t.Fatalf("Failed to insert test bid: %v", err)
		}
	}

	// Create a mock HTTP request with query parameters
	req, err := http.NewRequest("GET", "/find?collection=bid&filter={\"Description\":\"Second bid\"}", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	FindBidHandler(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Decode the response body into a slice of Bid structs
	var foundBids []structure.Bid
	if err := json.NewDecoder(rr.Body).Decode(&foundBids); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check if the found bids contain the expected bid
	var found bool
	expectedDescription := "Second bid"
	for _, bid := range foundBids {
		if bid.Description == expectedDescription {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected bid with description '%s' not found in retrieved bids", expectedDescription)
	}
}
