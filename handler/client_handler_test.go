package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Go-sumon/database"
	"Go-sumon/structure"
)

func TestCreateClientHandler(t *testing.T) {

    // Clear the user and client collections before running the test
    database.ClearCollection("user")
    database.ClearCollection("client")

    // Create a sample client payload
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
        Location: "New York",
    }
    clientJSON, err := json.Marshal(client)
    if err != nil {
        t.Fatal(err)
    }

    // Create a request with the sample payload
    req, err := http.NewRequest("POST", "/create-client", bytes.NewBuffer(clientJSON))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler function
    CreateClientHandler(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := `{"message":"Document created successfully"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }

    
}