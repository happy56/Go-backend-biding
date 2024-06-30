package database

import (
	"Go-sumon/structure"
	"context"
	"errors"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func UserCreate(userCollection string, document interface{}) error {

	initMongoClient(userCollection)

	// Check if the document is of type User
	user, ok := document.(*structure.User)
	if !ok {
		return errors.New("document is not of type User")
	}

	// Check if the phone number already exists
	filter := bson.M{"phonenumber": user.PhoneNumber}
	existingUsers := []*structure.User{} // Use a slice to hold multiple results
	err := Find(userCollection, filter, &existingUsers)
	if err != nil {
		return fmt.Errorf("failed to check phone number uniqueness: %v", err)
	}

	// If any existing users found, check for uniqueness
	if len(existingUsers) > 0 {
		return errors.New("cannot insert data because phone number already exists")
	}

	// Determine the current maximum UserID (if applicable)
	maxUserID := 0
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("failed to count existing users: %v", err)
	}
	maxUserID = int(count) + 1

	// Set the UserID for the new document
	v := reflect.ValueOf(document).Elem()
	userIDField := v.FieldByName("UserID")
	if userIDField.IsValid() && userIDField.CanSet() && maxUserID > 0 {
		userIDField.SetInt(int64(maxUserID))
	} else {
		return errors.New("UserID field not found or not settable")
	}

	// Call the create function passing the collection and document
	err = Create(userCollection, document)
	if err != nil {
		return err
	}

	return nil
}

func ClientCreate(document *structure.Client) error {
	// Call the UserCreate function to create the user document
	err := UserCreate("user", &document.User)
	if err != nil {
		return err
	}

	// If user type is "client", insert the user document into the client collection
	if document.User.UserType == "client" {
		// Call the create function passing the collection and document
		err = Create("client", document)
		if err != nil {
			return err
		}
	}

	return nil
}

func SPCreate(document *structure.ServiceProvider) error {
	// Call the UserCreate function to create the user document
	err := UserCreate("user", &document.User)
	if err != nil {
		return err
	}

	// If user type is "serviceProvider", insert the user document into the serviceProvider collection
	if document.User.UserType == "serviceProvider" {
		// Call the create function passing the collection and document
		err = Create("serviceProvider", document)
		if err != nil {
			return err
		}
	}

	return nil
}

func ClientUpdate(userID string, userData bson.M) error {
	// Update user collection
	if err := Update("user", userID, userData); err != nil {
		return err
	}

	// Update client collection
	if err := Update("client", userID, userData); err != nil {
		return err
	}

	return nil
}

func SPUpdate(userID string, userData bson.M) error {
	// Update user collection
	if err := Update("user", userID, userData); err != nil {
		return err
	}

	// Update serviceProvider collection
	if err := Update("serviceProvider", userID, userData); err != nil {
		return err
	}

	return nil
}
