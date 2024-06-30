package database

import (
	"Go-sumon/structure"
	"reflect"
	"testing"
)

func TestCreateSP(t *testing.T) {
    // Arrange: Clear the "user" and "serviceProvider" collections
    ClearCollection("serviceProvider")
    
    // Act: Define a service provider document to insert
    serviceProvider := structure.ServiceProvider{
        User: structure.User{
            // Manually setting other fields
            Name:        "Service Provider 1",
            PhoneNumber: "01711377008",
            NID:         "198426662698745",
            Birthdate:   "05-06-1984",
            FatherName:  "Father 1",
            MotherName:  "Mother 1",
            UserType:    "serviceProvider",
        },
        Skill:              "Plumbing",
        Location:           "City X",
        Education:          structure.Education{Level: "Bachelor", Institute: "University Y"},
        VerifiedByPorichoy: true,
        SPBalance:          structure.Balance{Amount: 1000},
    }

    // Call the SPCreate function to insert the service provider document
    if err := SPCreate(&serviceProvider); err != nil {
        t.Fatalf("Failed to insert service provider document: %v", err)
    }

    // Assert: Retrieve the inserted document from the "user" collection and compare with the expected values
    var insertedUser structure.User
    if err := Get("user", &insertedUser, serviceProvider.User.ID.Hex()); err != nil {
        t.Fatalf("Failed to retrieve inserted user document: %v", err)
    }

    // Check if the ID field of the inserted user document is set
    if insertedUser.ID.IsZero() {
        t.Error("Expected ID field to be set in the inserted user document")
    }

    // Assert: Retrieve the inserted document from the "serviceProvider" collection and compare with the expected values
    var insertedSP structure.ServiceProvider
    if err := Get("serviceProvider", &insertedSP, serviceProvider.User.ID.Hex()); err != nil {
        t.Fatalf("Failed to retrieve inserted service provider document: %v", err)
    }

    // Check if the ID field of the inserted service provider document is set
    if insertedSP.User.ID.IsZero() {
        t.Error("Expected ID field to be set in the inserted service provider document")
    }

    // Check if the inserted service provider matches the expected values
    expectedSP := serviceProvider
    if !reflect.DeepEqual(insertedSP, expectedSP) {
        t.Errorf("Unexpected service provider data, got %v, expected %v", insertedSP, expectedSP)
    }
}
