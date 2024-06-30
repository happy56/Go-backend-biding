package database

import (
	"Go-sumon/structure"
	"encoding/json"
	"fmt"
	"testing"
)

// func TestGetAllClient(t *testing.T) {
// 	// Arrange
// 	ClearCollection("client")

// 	// Insert some test data into the "client" collection
// 	clients := []structure.Client{}
// 	for i := 0; i < 2; i++ {
// 		id := primitive.NewObjectID() // Generate a new ObjectID
// 		client := structure.Client{
// 			ID:   id,
// 			User: structure.User{}, // Populate with necessary data
// 		}
// 		clients = append(clients, client)

// 		if err := Create("client", &client); err != nil {
// 			t.Fatalf("Failed to insert client document: %v", err)
// 		}
// 	}

// 	// Act: Call the GetAll function
// 	var resultClients []structure.Client
// 	if err := GetAll("client", &resultClients); err != nil {
// 		t.Fatalf("Failed to retrieve documents from the client collection: %v", err)
// 	}

// 	// Assert: Check if the retrieved documents match the inserted ones
// 	if len(resultClients) != len(clients) {
// 		t.Fatalf("Expected %d documents, got %d", len(clients), len(resultClients))
// 	}
// }

func TestCreateClient(t *testing.T) {
    // Arrange: Clear the "user" and "client" collections
    ClearCollection("user")
    ClearCollection("client")
    
    // Act: Define a client document to insert
    client := structure.Client{
        User: structure.User{
            // Manually setting other fields
            Name:        "Client 1",
            PhoneNumber: "01711377006",
            NID:         "198426662698745",
            Birthdate:   "05-06-1984",
            FatherName:  "Father 1",
            MotherName:  "Mother 1",
            UserType:    "client",
        },
        Location: "City X",
    }

    // Print the JSON value of the client document
    clientJSON, err := json.MarshalIndent(client, "", "    ")
    if err != nil {
        t.Fatalf("Failed to marshal client document to JSON: %v", err)
    }
    fmt.Println("Client JSON:", string(clientJSON))

    // Call the ClientCreate function to insert the client document
    if err := ClientCreate(&client); err != nil {
        t.Fatalf("Failed to insert client document: %v", err)
    }

    // Assert: Retrieve the inserted document from the "user" collection and compare with the expected values
    var insertedUser structure.User
    if err := Get("user", &insertedUser, client.User.ID.Hex()); err != nil {
        t.Fatalf("Failed to retrieve inserted user document: %v", err)
    }

    // Check if the ID field of the inserted user document is set
    if insertedUser.ID.IsZero() {
        t.Error("Expected ID field to be set in the inserted user document")
    }

    // Assert: Retrieve the inserted document from the "client" collection and compare with the expected values
    var insertedClient structure.Client
    if err := Get("client", &insertedClient, insertedUser.ID.Hex()); err != nil {
        t.Fatalf("Failed to retrieve inserted client document: %v", err)
    }

    // Check if the ID field of the inserted client document is set
    if insertedClient.User.ID.IsZero() {
        t.Error("Expected ID field to be set in the inserted client document")
    }
}


// func TestGetClient(t *testing.T) {
// 	// Arrange
// 	collectionName := "client"
// 	ClearCollection(collectionName)

// 	// Insert a client document
// 	expectedClient := structure.Client{
// 		ID:   primitive.NewObjectID(),
// 		User: structure.User{}, // Populate with necessary data
// 	}
// 	if err := Create(collectionName, &expectedClient); err != nil {
// 		t.Fatalf("Failed to insert test client document: %v", err)
// 	}

// 	// Act
// 	var resultClient structure.Client
// 	err := Get(collectionName, &resultClient, expectedClient.ID.Hex())

// 	// Assert
// 	if err != nil {
// 		t.Fatalf("Failed to retrieve client document: %v", err)
// 	}
// }

// func TestUpdateClient(t *testing.T) {
// 	// Arrange
// 	collectionName := "client"
// 	ClearCollection(collectionName)

// 	// Insert a client document
// 	expectedClient := structure.Client{
// 		ID:   primitive.NewObjectID(),
// 		User: structure.User{}, // Populate with necessary data
// 	}
// 	if err := Create(collectionName, &expectedClient); err != nil {
// 		t.Fatalf("Failed to insert test client document: %v", err)
// 	}

// 	// Define the update data
// 	updateData := bson.M{
// 		"user": bson.M{
// 			"name": "Updated Name",
// 		},
// 	}

// 	// Act: Update the client document
// 	if err := Update(collectionName, expectedClient.ID.Hex(), updateData); err != nil {
// 		t.Fatalf("Failed to update client document: %v", err)
// 	}

// 	// Retrieve the updated document
// 	var updatedClient structure.Client
// 	err := Get(collectionName, &updatedClient, expectedClient.ID.Hex())

// 	// Assert: Verify if the retrieved document matches the expected values
// 	if err != nil {
// 		t.Fatalf("Failed to retrieve updated client document: %v", err)
// 	}

// 	// Output the updated document with only the updated values
// 	fmt.Println("_id:", updatedClient.ID.Hex())
// 	fmt.Println("name:", updatedClient.User.Name)
// }

// func TestDeleteClient(t *testing.T) {
// 	// Arrange
// 	collectionName := "client"
// 	ClearCollection(collectionName)

// 	// Insert three client documents
// 	var clients []structure.Client
// 	for i := 0; i < 3; i++ {
// 		client := structure.Client{
// 			ID:   primitive.NewObjectID(),
// 			User: structure.User{}, // Populate with necessary data
// 		}
// 		clients = append(clients, client)
// 		if err := Create(collectionName, &client); err != nil {
// 			t.Fatalf("Failed to insert test client document %d: %v", i+1, err)
// 		}
// 	}

// 	// Delete one client document
// 	err := Delete(collectionName, clients[0].ID.Hex())

// 	// Assert
// 	if err != nil {
// 		t.Fatalf("Failed to delete client document: %v", err)
// 	}

// 	// Attempt to retrieve the deleted document
// 	var resultClient structure.Client
// 	err = Get(collectionName, &resultClient, clients[0].ID.Hex())

// 	// Assert: Check if the document is not found
// 	expectedErrMsg := fmt.Sprintf("document with ID %s not found in collection %s", clients[0].ID.Hex(), collectionName)
// 	if err == nil || err.Error() != expectedErrMsg {
// 		t.Errorf("Expected error '%s', got '%v'", expectedErrMsg, err)
// 	}
// }

// func TestFindClient(t *testing.T) {
// 	// Arrange
// 	collectionName := "client"
// 	ClearCollection(collectionName)

// 	// Insert two client documents
// 	expectedClient1 := structure.Client{
// 		ID:   primitive.NewObjectID(),
// 		User: structure.User{}, // Populate with necessary data
// 	}
// 	expectedClient2 := structure.Client{
// 		ID:   primitive.NewObjectID(),
// 		User: structure.User{}, // Populate with necessary data
// 	}
// 	if err := Create(collectionName, &expectedClient1); err != nil {
// 		t.Fatalf("Failed to insert test client document 1: %v", err)
// 	}
// 	if err := Create(collectionName, &expectedClient2); err != nil {
// 		t.Fatalf("Failed to insert test client document 2: %v", err)
// 	}

// 	// Act
// 	var resultClients []structure.Client
// 	err := Find(collectionName, bson.M{}, &resultClients)

// 	// Assert
// 	if err != nil {
// 		t.Fatalf("Failed to retrieve client documents: %v", err)
// 	}
// 	if len(resultClients) != 2 {
// 		t.Errorf("Expected 2 client documents, got %d", len(resultClients))
// 	}
// }
