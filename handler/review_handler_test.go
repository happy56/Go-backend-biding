package handler

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"net/http"
	"net/http/httptest"
	"net/url"

	"testing"

	"Go-sumon/database"
	"Go-sumon/structure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllReviewHandler(t *testing.T) {
	// Clear the "review" collection before running the test
	database.ClearCollection("review")

	// Insert two review documents for testing
	review1 := structure.Review{
		ID:            primitive.NewObjectID(),
		Review:        "Great service",
		Timelines:     4.5,
		Quality:       4.2,
		Communication: 4.8,
		Behavior:      4.6,
	}
	review2 := structure.Review{
		ID:            primitive.NewObjectID(),
		Review:        "Excellent service",
		Timelines:     4.7,
		Quality:       4.4,
		Communication: 4.9,
		Behavior:      4.7,
	}
	if err := database.Create("review", &review1); err != nil {
		t.Fatalf("Failed to insert test review document 1: %v", err)
	}
	if err := database.Create("review", &review2); err != nil {
		t.Fatalf("Failed to insert test review document 2: %v", err)
	}

	// Create a new HTTP GET request
	req := httptest.NewRequest("GET", "/review", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	GetAllReviewHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	// Decode the response body into a slice of Review objects
	var reviews []structure.Review
	if err := json.NewDecoder(rr.Body).Decode(&reviews); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check if the correct number of reviews was retrieved
	if len(reviews) != 2 {
		t.Errorf("unexpected number of reviews retrieved: got %d, want 2", len(reviews))
	}

	// Check the content of the first review
	if !reflect.DeepEqual(reviews[0], review1) {
		t.Errorf("unexpected content of the first review: got %+v, want %+v", reviews[0], review1)
	}

	// Check the content of the second review
	if !reflect.DeepEqual(reviews[1], review2) {
		t.Errorf("unexpected content of the second review: got %+v, want %+v", reviews[1], review2)
	}
}

func TestCreateReviewHandler(t *testing.T) {

	database.ClearCollection("review")
	// Create a new HTTP request
	requestBody := strings.NewReader(`{"Review": "Great service", "Timelines": 4.5, "Quality": 4.2, "Communication": 4.8, "Behavior": 4.6}`)
	req := httptest.NewRequest("POST", "/review", requestBody)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	CreateReviewHandler(rr, req)

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
	var createdReview structure.Review
	err := json.NewDecoder(rr.Body).Decode(&createdReview)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check if the review was created successfully
	if createdReview.ID.IsZero() {
		t.Error("review ID was not set")
	}
	if createdReview.Review != "Great service" {
		t.Errorf("unexpected review name: got %s, want Great service", createdReview.Review)
	}
	// Check other fields as needed
}

func TestGetReviewHandler(t *testing.T) {
	// Clear the "review" collection before running the test
	database.ClearCollection("review")

	// Create mock review documents for testing
	mockReview1 := structure.Review{
		ID:            primitive.NewObjectID(),
		Review:        "Great service 1",
		Timelines:     4.5,
		Quality:       4.2,
		Communication: 4.8,
		Behavior:      4.6,
	}

	mockReview2 := structure.Review{
		ID:            primitive.NewObjectID(),
		Review:        "Great service 2",
		Timelines:     4.3,
		Quality:       4.1,
		Communication: 4.7,
		Behavior:      4.5,
	}

	// Insert the mock review documents into the database
	if err := database.Create("review", &mockReview1); err != nil {
		t.Fatalf("Failed to insert mock review document 1: %v", err)
	}
	if err := database.Create("review", &mockReview2); err != nil {
		t.Fatalf("Failed to insert mock review document 2: %v", err)
	}

	// Create a new HTTP GET request with the ID of the second mock review in the URL query
	req, err := http.NewRequest("GET", "/review?id="+mockReview2.ID.Hex(), nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and response recorder
	GetReviewHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a slice of Review objects
	var retrievedReviews []structure.Review
	if err := json.NewDecoder(rr.Body).Decode(&retrievedReviews); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
}

func TestUpdateReviewHandler(t *testing.T) {
	// Clear the "review" collection before running the test
	database.ClearCollection("review")

	// Insert a review document for testing
	expectedReview := structure.Review{
		Review:        "Great service", // <-- This needs to be "Updated service"
		Timelines:     4.5,
		Quality:       4.2,
		Communication: 4.8,
		Behavior:      4.6,
	}
	if err := database.Create("review", &expectedReview); err != nil {
		t.Fatalf("Failed to insert test review document: %v", err)
	}

	// Define the update data
	updateData := map[string]interface{}{
		"review":        "Updated service",
		"timelines":     4.3,
		"quality":       4.1,
		"communication": 4.7,
		"behavior":      4.5,
	}

	// Marshal the update data to JSON
	updateBody, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("Failed to marshal update data: %v", err)
	}

	// Convert the ObjectID to a string
	idString := expectedReview.ID.Hex()

	// Create a PUT request with the update data
	req, err := http.NewRequest("PUT", "/review?id="+idString, bytes.NewBuffer(updateBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Execute the handler function
	UpdateReviewHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Retrieve the updated document from the database
	var updatedReview structure.Review
	err = database.Get("review", &updatedReview, idString)
	if err != nil {
		t.Fatalf("Failed to retrieve updated review document: %v", err)
	}

}

func TestDeleteReviewHandler(t *testing.T) {
	// Clear the "review" collection before running the test
	database.ClearCollection("review")

	// Insert a review document for testing
	expectedReview1 := structure.Review{
		Review:        "Test review 1",
		Timelines:     4.5,
		Quality:       4.2,
		Communication: 4.8,
		Behavior:      4.6,
	}
	if err := database.Create("review", &expectedReview1); err != nil {
		t.Fatalf("Failed to insert test review document 1: %v", err)
	}

	// Insert another review document for testing
	expectedReview2 := structure.Review{
		Review:        "Test review 2",
		Timelines:     4.1,
		Quality:       4.3,
		Communication: 4.6,
		Behavior:      4.8,
	}
	if err := database.Create("review", &expectedReview2); err != nil {
		t.Fatalf("Failed to insert test review document 2: %v", err)
	}

	// Convert the ObjectID of the second document to a string
	id := expectedReview2.ID.Hex()

	// Create a DELETE request with the review ID of the second document
	req, err := http.NewRequest("DELETE", "/review?id="+id, nil)
	if err != nil {
		t.Fatalf("Failed to create delete request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Execute the handler function
	DeleteReviewHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the second document still exists in the database
	var remainingReview structure.Review
	err = database.Get("review", &remainingReview, id)
	if err == nil {
		t.Errorf("Second document still exists in the database after deletion")
	}
}

func TestFindReviewHandler(t *testing.T) {
	// Clear the "review" collection before running the test
	database.ClearCollection("review")

	// Insert test reviews into the database
	testReviews := []structure.Review{
		{Review: "First review", Timelines: 4.5, Quality: 4.2, Communication: 4.8, Behavior: 4.6},
		{Review: "Second review", Timelines: 4.1, Quality: 4.3, Communication: 4.6, Behavior: 4.8},
	}

	for _, review := range testReviews {
		if err := database.Create("review", &review); err != nil {
			t.Fatalf("Failed to insert test review: %v", err)
		}
	}

	// Create a mock HTTP request with query parameters
	query := url.Values{}
	query.Set("collection", "review")
	query.Set("filter", `{"Review":"Second review"}`)
	req, err := http.NewRequest("GET", "/find?"+query.Encode(), nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	FindReviewHandler(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Decode the response body into a slice of Review structs
	var foundReviews []structure.Review
	if err := json.NewDecoder(rr.Body).Decode(&foundReviews); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check if the found reviews match the expected review
	var expectedReview structure.Review
	for _, review := range testReviews {
		if review.Review == "Second review" {
			expectedReview = review
			break
		}
	}

	if !reflect.DeepEqual(foundReviews, []structure.Review{expectedReview}) {
		t.Errorf("Found review does not match expected review: got %v, want %v", foundReviews, []structure.Review{expectedReview})
	}
}
