package database

import (
	"Go-sumon/structure"
	"context"
	"fmt"

	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllUser(t *testing.T) {
	// Arrange: Clear the "user" collection
	ClearCollection("user")

	// Get the count of existing users
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		t.Fatalf("Failed to count existing users: %v", err)
	}

	// Insert some test data into the "user" collection
	users := []structure.User{}
	for i := int64(0); i < 10; i++ {
		id := primitive.NewObjectID() // Generate a new ObjectID
		user := structure.User{
			ID:          id, // Assign the ObjectID directly
			Name:        fmt.Sprintf("User %d", count+i+1),
			PhoneNumber: fmt.Sprintf("123456789%d", count+i),
			NID:         fmt.Sprintf("NID%d", count+i+1),
			Birthdate:   fmt.Sprintf("200%d-01-01", count+i+1),
			FatherName:  fmt.Sprintf("Father %d", count+i+1),
			MotherName:  fmt.Sprintf("Mother %d", count+i+1),
			UserType:    structure.UserTypeClient, // Set the userType
		}
		// Set the UserID of the new user
		user.UserID = int(count + i + 1)
		users = append(users, user)

		if err := UserCreate("user", &user); err != nil {
			t.Fatalf("Failed to insert user document: %v", err)
		}
	}

	// Act: Call the GetAll function to get all documents from "user" collection
	var resultUsers []structure.User
	if err := GetAll("user", &resultUsers); err != nil {
		t.Fatalf("Failed to retrieve documents from the user collection: %v", err)
	}

	// Assert: Check if the retrieved documents match the inserted ones
	if len(resultUsers) != len(users) {
		t.Fatalf("Expected %d documents, got %d", len(users), len(resultUsers))
	}
	for i := range users {
		if !reflect.DeepEqual(resultUsers[i], users[i]) {
			t.Errorf("Retrieved users do not match the inserted ones")
		}
	}
}

func TestCreateUser(t *testing.T) {
	// Arrange: Clear the "user" and "serviceProvider" collections
	ClearCollection("user")
	ClearCollection("client")
	ClearCollection("serviceProvider")

	// Define a client user document to insert
	clientUser := structure.User{
		Name:        "Client User",
		PhoneNumber: "0123456789",
		NID:         "NID1",
		Birthdate:   "2001-01-01",
		FatherName:  "Client Father",
		MotherName:  "Client Mother",
		UserType:    structure.UserTypeClient, // Set the userType
	}

	// Define a service provider user document to insert
	serviceProviderUser := structure.User{
		Name:        "Service Provider User",
		PhoneNumber: "0987654321",
		NID:         "NID2",
		Birthdate:   "1990-01-01",
		FatherName:  "SP Father",
		MotherName:  "SP Mother",
		UserType:    structure.UserTypeServiceProvider, // Set the userType
	}

	// Act: Call the UserCreate function to insert the client user document
	if err := UserCreate("user", &clientUser); err != nil {
		t.Fatalf("Failed to insert client user document: %v", err)
	}

	// Act: Call the UserCreate function to insert the service provider user document
	if err := UserCreate("user", &serviceProviderUser); err != nil {
		t.Fatalf("Failed to insert service provider user document: %v", err)
	}

	// Assert: Retrieve the inserted client user document and compare with the expected values
	var insertedClientUser structure.User
	if err := Get("user", &insertedClientUser, clientUser.ID.Hex()); err != nil {
		t.Fatalf("Failed to retrieve inserted client user document: %v", err)
	}

	// Check if the ID field of the inserted client user document is set
	if insertedClientUser.ID.IsZero() {
		t.Error("Expected ID field to be set in the inserted client user document")
	}

	// Check if the inserted client user matches the expected values
	expectedClientUser := clientUser
	if !reflect.DeepEqual(insertedClientUser, expectedClientUser) {
		t.Errorf("Unexpected client user data, got %v, expected %v", insertedClientUser, expectedClientUser)
	}

	// Assert: Retrieve the inserted service provider user document and compare with the expected values
	var insertedServiceProviderUser structure.User
	if err := Get("user", &insertedServiceProviderUser, serviceProviderUser.ID.Hex()); err != nil {
		t.Fatalf("Failed to retrieve inserted service provider user document: %v", err)
	}

	// Check if the ID field of the inserted service provider user document is set
	if insertedServiceProviderUser.ID.IsZero() {
		t.Error("Expected ID field to be set in the inserted service provider user document")
	}

	// Check if the inserted service provider user matches the expected values
	expectedServiceProviderUser := serviceProviderUser
	if !reflect.DeepEqual(insertedServiceProviderUser, expectedServiceProviderUser) {
		t.Errorf("Unexpected service provider user data, got %v, expected %v", insertedServiceProviderUser, expectedServiceProviderUser)
	}
}

func TestGetUser(t *testing.T) {
	// Arrange
	collectionName := "user"
	ClearCollection(collectionName)

	// Insert two user documents
	expectedUser1 := structure.User{
		Name:        "User One",
		PhoneNumber: "1234567890",
		NID:         "NID123456",
		Birthdate:   "2000-01-01",
		FatherName:  "Father One",
		MotherName:  "Mother One",
		UserType:    structure.UserTypeClient,
	}
	expectedUser2 := structure.User{
		Name:        "User Two",
		PhoneNumber: "0987654321",
		NID:         "NID654321",
		Birthdate:   "2001-01-01",
		FatherName:  "Father Two",
		MotherName:  "Mother Two",
		UserType:    structure.UserTypeServiceProvider,
	}
	if err := UserCreate(collectionName, &expectedUser1); err != nil {
		t.Fatalf("Failed to insert test user document 1: %v", err)
	}
	if err := UserCreate(collectionName, &expectedUser2); err != nil {
		t.Fatalf("Failed to insert test user document 2: %v", err)
	}

	// Act
	var resultUser structure.User
	err := Get(collectionName, &resultUser, expectedUser1.ID.Hex())

	// Assert
	if err != nil {
		t.Fatalf("Failed to retrieve user document: %v", err)
	}
	if !reflect.DeepEqual(resultUser, expectedUser1) {
		t.Errorf("Retrieved user document does not match expected document")
	}
}

func TestUpdateUser(t *testing.T) {
	// Arrange
	collectionName := "user"
	ClearCollection(collectionName)

	// Insert a user document
	expectedUser := structure.User{
		Name:        "Initial User",
		PhoneNumber: "1234567890",
		NID:         "NID123456",
		Birthdate:   "2000-01-01",
		FatherName:  "Initial Father",
		MotherName:  "Initial Mother",
		UserType:    structure.UserTypeClient, // Assign UserType
	}
	if err := UserCreate(collectionName, &expectedUser); err != nil {
		t.Fatalf("Failed to insert test user document: %v", err)
	}

	// Define the update data
	updateData := bson.M{
		"name":       "Updated User",
		"nid":        "NID654321",
		"birthdate":  "2000-01-02",
		"fathername": "Updated Father",
		"mothername": "Updated Mother",
	}

	// Act: Update the user document
	if err := Update(collectionName, expectedUser.ID.Hex(), updateData); err != nil {
		t.Fatalf("Failed to update user document: %v", err)
	}

	// Retrieve the updated document
	var updatedUser structure.User
	err := Get(collectionName, &updatedUser, expectedUser.ID.Hex())

	// Assert: Verify if the retrieved document matches the expected values
	if err != nil {
		t.Fatalf("Failed to retrieve updated user document: %v", err)
	}

	// Output the updated fields only
	fmt.Println("ID:", updatedUser.ID.Hex())
	fmt.Println("name:", updatedUser.Name)
	fmt.Println("phoneNumber:", updatedUser.PhoneNumber)
	fmt.Println("nid:", updatedUser.NID)
	fmt.Println("birthdate:", updatedUser.Birthdate)
	fmt.Println("fatherName:", updatedUser.FatherName)
	fmt.Println("motherName:", updatedUser.MotherName)

}

func TestDeleteUser(t *testing.T) {
	// Arrange
	collectionName := "user"
	ClearCollection(collectionName)

	// Insert three user documents
	expectedUser1 := structure.User{
		ID:          primitive.NewObjectID(),
		Name:        "User One",
		PhoneNumber: "1234567890",
		NID:         "NID123456",
		Birthdate:   "2000-01-01",
		FatherName:  "Father One",
		MotherName:  "Mother One",
		UserType:    structure.UserTypeClient, // User type included
	}
	expectedUser2 := structure.User{
		ID:          primitive.NewObjectID(),
		Name:        "User Two",
		PhoneNumber: "0987654321",
		NID:         "NID654321",
		Birthdate:   "2001-01-01",
		FatherName:  "Father Two",
		MotherName:  "Mother Two",
		UserType:    structure.UserTypeClient, // User type included
	}
	expectedUser3 := structure.User{
		ID:          primitive.NewObjectID(),
		Name:        "User Three",
		PhoneNumber: "0123456789",
		NID:         "NID789012",
		Birthdate:   "2002-01-01",
		FatherName:  "Father Three",
		MotherName:  "Mother Three",
		UserType:    structure.UserTypeServiceProvider, // User type included
	}
	if err := UserCreate(collectionName, &expectedUser1); err != nil {
		t.Fatalf("Failed to insert test user document 1: %v", err)
	}
	if err := UserCreate(collectionName, &expectedUser2); err != nil {
		t.Fatalf("Failed to insert test user document 2: %v", err)
	}
	if err := UserCreate(collectionName, &expectedUser3); err != nil {
		t.Fatalf("Failed to insert test user document 3: %v", err)
	}

	// Act: Call the Delete function for the second user document
	err := Delete(collectionName, expectedUser2.ID.Hex())

	// Assert
	if err != nil {
		t.Fatalf("Failed to delete user document: %v", err)
	}

	// Retrieve the remaining user documents
	var resultUsers []structure.User
	err = GetAll(collectionName, &resultUsers)
	if err != nil {
		t.Fatalf("Failed to retrieve user documents: %v", err)
	}

	// Assert: Verify the remaining documents and their UserID values
	if len(resultUsers) != 2 {
		t.Fatalf("Expected 2 remaining user documents, got %d", len(resultUsers))
	}

	// Check if the remaining documents match the expected ones
	if !reflect.DeepEqual(resultUsers[0], expectedUser1) {
		t.Errorf("Retrieved user document does not match expected document")
	}
	if !reflect.DeepEqual(resultUsers[1], expectedUser3) {
		t.Errorf("Retrieved user document does not match expected document")
	}
}

func TestFindUser(t *testing.T) {
	// Arrange
	collectionName := "user"
	ClearCollection(collectionName)

	// Insert two user documents
	expectedUser1 := structure.User{
		ID:          primitive.NewObjectID(),
		Name:        "User One",
		PhoneNumber: "1234567890",
		NID:         "NID123456",
		Birthdate:   "2000-01-01",
		FatherName:  "Father One",
		MotherName:  "Mother One",
		UserType:    structure.UserTypeClient,
	}
	expectedUser2 := structure.User{
		ID:          primitive.NewObjectID(),
		Name:        "User Two",
		PhoneNumber: "0987654321",
		NID:         "NID654321",
		Birthdate:   "2001-01-01",
		FatherName:  "Father Two",
		MotherName:  "Mother Two",
		UserType:    structure.UserTypeServiceProvider,
	}
	if err := UserCreate(collectionName, &expectedUser1); err != nil {
		t.Fatalf("Failed to insert test user document 1: %v", err)
	}
	if err := UserCreate(collectionName, &expectedUser2); err != nil {
		t.Fatalf("Failed to insert test user document 2: %v", err)
	}

	// Act
	var resultUsers []structure.User
	err := Find(collectionName, bson.M{}, &resultUsers)

	// Assert
	if err != nil {
		t.Fatalf("Failed to retrieve user documents: %v", err)
	}
	if len(resultUsers) != 2 {
		t.Errorf("Expected 2 user documents, got %d", len(resultUsers))
	}
}
