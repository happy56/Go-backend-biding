package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"Go-sumon/database"

	"go.mongodb.org/mongo-driver/bson"
	
)

func GenericGetAllHandler(w http.ResponseWriter, r *http.Request, collectionName string, result interface{}) {
	// Set Access-Control-Allow-Origin header to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call the GetAll function to retrieve all items from the specified collection in the database
	err := database.GetAll(collectionName, result)
	if err != nil {
		http.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 (OK) before writing the response body
	w.WriteHeader(http.StatusOK)

	// Respond with the retrieved items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GenericCreateHandler(w http.ResponseWriter, r *http.Request, collectionName string, document interface{}) {
	// Set Access-Control-Allow-Origin header to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body into the provided document interface
	err := json.NewDecoder(r.Body).Decode(&document)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Call the provided Create function to insert the document into the specified collection
	err = database.Create(collectionName, document)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create document in collection %s: %v", collectionName, err), http.StatusInternalServerError)
		return
	}

	// Set the status code to 201 (Created) before writing the response body
	w.WriteHeader(http.StatusCreated)

	// Respond with the newly created document
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

func GenericGetHandler(w http.ResponseWriter, r *http.Request, collectionName string) {
    // Set Access-Control-Allow-Origin header to allow cross-origin requests
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // Check if the request method is GET
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse request parameters
    params := r.URL.Query()
    id := params.Get("id")
    if id == "" {
        http.Error(w, "Missing ID parameter", http.StatusBadRequest)
        return
    }

    // Call the Get function to retrieve the document
    var result interface{}
    err := database.Get(collectionName, &result, id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get document: %v", err), http.StatusInternalServerError)
        return
    }

    // Set the status code to 200 (OK) before writing the response body
    w.WriteHeader(http.StatusOK)

    // Respond with the retrieved document
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func GenericUpdateHandler(w http.ResponseWriter, r *http.Request, collectionName string) {
	// Parse request parameters
	params := r.URL.Query()
	id := params.Get("id")

	// Decode request body
	var updateData bson.M
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Update document in the database
	err = database.Update(collectionName, id, updateData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update document: %v", err), http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
}

func GenericDeleteHandler(w http.ResponseWriter, r *http.Request, collectionName string) {
	// Set Access-Control-Allow-Origin header to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the ID from the URL query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	// Print the ID being used for deletion
	fmt.Printf("Attempting to delete document with ID: %s from collection: %s\n", id, collectionName)

	// Call the delete function to delete the document from the specified collection
	err := database.Delete(collectionName, id)
	if err != nil {
		log.Printf("Error deleting document: %v\n", err) // Log the error

		// Print the ID being used for deletion
		fmt.Printf("Failed to delete document with ID: %s\n", id)

		http.Error(w, fmt.Sprintf("Failed to delete document: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 (OK) before writing the response body
	w.WriteHeader(http.StatusOK)

	// Write a simple success message
	_, err = w.Write([]byte("Document deleted successfully"))
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}

func GenericFindHandler(w http.ResponseWriter, r *http.Request, collectionName string) {
	// Set Access-Control-Allow-Origin header to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the filter from the URL query parameters
	filterParam := r.URL.Query().Get("filter")
	if filterParam == "" {
		http.Error(w, "Filter not provided", http.StatusBadRequest)
		return
	}

	// Convert the filter string to a map[string]interface{}
	var filter interface{}
	err := json.Unmarshal([]byte(filterParam), &filter)
	if err != nil {
		http.Error(w, "Invalid filter format", http.StatusBadRequest)
		return
	}

	// Define a variable to hold the result of the find operation
	// Assuming result is of type interface{}, you can modify this based on your implementation
	var result interface{}

	// Call the Find function to retrieve documents from the specified collection
	err = database.Find(collectionName, filter, &result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find documents: %v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the result to JSON
	responseBody, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal response: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write(responseBody)
}
