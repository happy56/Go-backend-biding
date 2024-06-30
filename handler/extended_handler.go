package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/http"

	"Go-sumon/database"
	"Go-sumon/structure"
)

func UserCreateHandler(w http.ResponseWriter, r *http.Request, collectionName string) {
	// Set Access-Control-Allow-Origin header to allow cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body to get the document data
	var document structure.User
	err := json.NewDecoder(r.Body).Decode(&document)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Call the appropriate create function to create the document
	err = database.UserCreate(collectionName, document)
	if err != nil {
		if strings.Contains(err.Error(), "phone number already exists") {
			http.Error(w, "Phone number already exists", http.StatusConflict)
		} else {
			http.Error(w, fmt.Sprintf("Failed to create document: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Set the status code to 201 (Created)
	w.WriteHeader(http.StatusCreated)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the result (assuming result is defined somewhere in your code)
	json.NewEncoder(w).Encode(map[string]string{"message": "Document created successfully"})
}


func ClientCreateHandler(w http.ResponseWriter, r *http.Request, userCollectionName string, clientCollectionName string) {
    // Set content type header
    w.Header().Set("Content-Type", "application/json")

    // Parse the request body into a structure.Client object
    var client structure.Client
    err := json.NewDecoder(r.Body).Decode(&client)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Call the UserCreate function to create the user document
    err = database.UserCreate(userCollectionName, &client.User)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // If user type is "client", insert the user document into the client collection
    if client.User.UserType == "client" {
        // Call the Create function passing the collection and document
        err = database.Create(clientCollectionName, &client)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Encode the response
    response := map[string]string{"message": "Document created successfully"}
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func SPCreateHandler(w http.ResponseWriter, r *http.Request, userCollectionName string, SPCollectionName string) {
    // Set content type header
    w.Header().Set("Content-Type", "application/json")

    // Parse the request body into a structure.Client object
    var serviceProvider structure.ServiceProvider
    err := json.NewDecoder(r.Body).Decode(&serviceProvider)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Call the UserCreate function to create the user document
    err = database.UserCreate(userCollectionName, &serviceProvider.User)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // If user type is "client", insert the user document into the client collection
    if serviceProvider.User.UserType == "serviceProvider" {
        // Call the Create function passing the collection and document
        err = database.Create(SPCollectionName, &serviceProvider)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Encode the response
    response := map[string]string{"message": "Document created successfully"}
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}