package database

import (
	"Go-sumon/structure"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllJob(t *testing.T) {
	// Arrange
	ClearCollection("job")

	// Insert some test data into the "job" collection
	jobs := []structure.Job{}
	for i := 0; i < 2; i++ {
		id := primitive.NewObjectID() // Generate a new ObjectID
		job := structure.Job{
			ID:               id,
			Title:            fmt.Sprintf("Test Job %d", i+1),
			Posted:           time.Now(),
			Budget:           "$1000",
			Description:      "Test job description",
			Clients:          structure.Client{},             // Populate with necessary data
			ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
			Status:           structure.StatusPending,        // Use constant values directly
			JobStatus:        structure.JobStatusJobPosted,   // Use constant values directly
			//StatusChange:     structure.StatusChangePostingTime, // Use constant values directly
			Point:            []structure.Point{},            // Populate with necessary data
			QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
			Bid:              []structure.Bid{},              // Populate with necessary data
			Review:           structure.Review{},           // Populate with necessary data
		}
		jobs = append(jobs, job)

		if err := Create("job", &job); err != nil {
			t.Fatalf("Failed to insert job document: %v", err)
		}
	}

	// Act: Call the GetAll function
	var resultJobs []structure.Job
	if err := GetAll("job", &resultJobs); err != nil {
		t.Fatalf("Failed to retrieve documents from the job collection: %v", err)
	}

	// Assert: Check if the retrieved documents match the inserted ones
	if len(resultJobs) != len(jobs) {
		t.Fatalf("Expected %d documents, got %d", len(jobs), len(resultJobs))
	}
	
}


func TestCreateJob(t *testing.T) {
	// Arrange
	ClearCollection("job")

	// Define a job document to insert
	job := structure.Job{
		ID:               primitive.NewObjectID(),
		Title:            "Test Job",
		Posted:           time.Now(),
		Budget:           "$1000",
		Description:      "Test job description",
		Clients:          structure.Client{},             // Populate with necessary data
		ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
		Status:           structure.StatusPending,        // Example value for Status
		JobStatus:        structure.JobStatusJobPosted,  // Example value for JobStatus
		//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
		Point:            []structure.Point{},            // Populate with necessary data
		QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
		Bid:              []structure.Bid{},              // Populate with necessary data
		Review:           structure.Review{},           // Populate with necessary data
	}

	// Act: Insert the job document
	if err := Create("job", &job); err != nil {
		t.Fatalf("Failed to insert job document: %v", err)
	}

	// Assert: Retrieve the inserted document and compare with the expected values
	var insertedJob structure.Job
	if err := Get("job", &insertedJob, job.ID.Hex()); err != nil {
		t.Fatalf("Failed to retrieve inserted document: %v", err)
	}

	// Check if the ID field of the inserted document is set
	if insertedJob.ID.IsZero() {
		t.Error("Expected ID field to be set in the inserted document")
	}

}

func TestGetJob(t *testing.T) {
	// Arrange
	collectionName := "job"
	ClearCollection(collectionName)

	// Insert a job document
	expectedJob := structure.Job{
		ID:               primitive.NewObjectID(),
		Title:            "Test Job",
		Posted:           time.Now(),
		Budget:           "$1000",
		Description:      "Test job description",
		Clients:          structure.Client{},             // Populate with necessary data
		ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
		Status:           structure.StatusPending,        // Example value for Status
		JobStatus:        structure.JobStatusJobPosted,  // Example value for JobStatus
		//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
		Point:            []structure.Point{},            // Populate with necessary data
		QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
		Bid:              []structure.Bid{},              // Populate with necessary data
		Review:           structure.Review{},           // Populate with necessary data
	}
	if err := Create(collectionName, &expectedJob); err != nil {
		t.Fatalf("Failed to insert test job document: %v", err)
	}

	// Act
	var resultJob structure.Job
	err := Get(collectionName, &resultJob, expectedJob.ID.Hex())

	// Assert
	if err != nil {
		t.Fatalf("Failed to retrieve job document: %v", err)
	}
	
}


func TestUpdateJob(t *testing.T) {
	// Arrange
	collectionName := "job"
	ClearCollection(collectionName)

	// Insert a job document
	expectedJob := structure.Job{
		ID:               primitive.NewObjectID(),
		Title:            "Test Job",
		Posted:           time.Now(),
		Budget:           "$1000",
		Description:      "Test job description",
		Clients:          structure.Client{},             // Populate with necessary data
		ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
		Status:           structure.StatusPending,        // Example value for Status
		JobStatus:        structure.JobStatusJobPosted,  // Example value for JobStatus
		//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
		Point:            []structure.Point{},            // Populate with necessary data
		QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
		Bid:              []structure.Bid{},              // Populate with necessary data
		Review:           structure.Review{},           // Populate with necessary data
	}
	if err := Create(collectionName, &expectedJob); err != nil {
		t.Fatalf("Failed to insert test job document: %v", err)
	}

	// Define the update data
	updateData := bson.M{
		"title":       "Updated Test Job",
		"description": "Updated job description",
		"budget":      "$2000",
	}

	// Act: Update the job document
	if err := Update(collectionName, expectedJob.ID.Hex(), updateData); err != nil {
		t.Fatalf("Failed to update job document: %v", err)
	}

	// Retrieve the updated document
	var updatedJob structure.Job
	err := Get(collectionName, &updatedJob, expectedJob.ID.Hex())

	// Assert: Verify if the retrieved document matches the expected values
	if err != nil {
		t.Fatalf("Failed to retrieve updated job document: %v", err)
	}

	// Output the updated document with only the updated values
	fmt.Println("_id:", updatedJob.ID.Hex())
	fmt.Println("title:", updatedJob.Title)
	fmt.Println("description:", updatedJob.Description)
	fmt.Println("budget:", updatedJob.Budget)
}

func TestDeleteJob(t *testing.T) {
	// Arrange
	collectionName := "job"
	ClearCollection(collectionName)

	// Insert three job documents
	var jobs []structure.Job
	for i := 0; i < 3; i++ {
		job := structure.Job{
			ID:               primitive.NewObjectID(),
			Title:            fmt.Sprintf("Test Job %d", i+1),
			Posted:           time.Now(),
			Budget:           fmt.Sprintf("$%d000", i+1),
			Description:      fmt.Sprintf("Test job description %d", i+1),
			Clients:          structure.Client{},             // Populate with necessary data
			ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
			Status:           structure.StatusPending,        // Example value for Status
			JobStatus:        structure.JobStatusJobPosted,  // Example value for JobStatus
			//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
			Point:            []structure.Point{},            // Populate with necessary data
			QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
			Bid:              []structure.Bid{},              // Populate with necessary data
			Review:           structure.Review{},           // Populate with necessary data
		}
		jobs = append(jobs, job)
		if err := Create(collectionName, &job); err != nil {
			t.Fatalf("Failed to insert test job document %d: %v", i+1, err)
		}
	}

	// Delete one job document
	err := Delete(collectionName, jobs[0].ID.Hex())

	// Assert
	if err != nil {
		t.Fatalf("Failed to delete job document: %v", err)
	}

	// Attempt to retrieve the deleted document
	var resultJob structure.Job
	err = Get(collectionName, &resultJob, jobs[0].ID.Hex())

	// Assert: Check if the document is not found
	expectedErrMsg := fmt.Sprintf("document with ID %s not found in collection %s", jobs[0].ID.Hex(), collectionName)
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error '%s', got '%v'", expectedErrMsg, err)
	}
}

func TestFindJob(t *testing.T) {
	// Arrange
	collectionName := "job"
	ClearCollection(collectionName)

	// Insert two job documents
	expectedJob1 := structure.Job{
		ID:               primitive.NewObjectID(),
		Title:            "Test Job 1",
		Posted:           time.Now(),
		Budget:           "$1000",
		Description:      "Test job description 1",
		Clients:          structure.Client{},             // Populate with necessary data
		ServiceProviders: structure.ServiceProvider{},    // Populate with necessary data
		Status:           structure.StatusPending,        // Example value for Status
		JobStatus:        structure.JobStatusJobPosted,  // Example value for JobStatus
		//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
		Point:            []structure.Point{},            // Populate with necessary data
		QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
		Bid:              []structure.Bid{},              // Populate with necessary data
		Review:           structure.Review{},           // Populate with necessary data
	}
	expectedJob2 := structure.Job{
		ID:               primitive.NewObjectID(),
		Title:            "Test Job 2",
		Posted:           time.Now(),
		Budget:           "$2000",
		Description:      "Test job description 2",
		Clients:          structure.Client{},               // Populate with necessary data
		ServiceProviders: structure.ServiceProvider{},      // Populate with necessary data
		Status:           structure.StatusPending,          // Example value for Status
		JobStatus:        structure.JobStatusJobPosted,    // Example value for JobStatus
		//StatusChange:     structure.StatusChangePostingTime, // Example value for StatusChange
		Point:            []structure.Point{},            // Populate with necessary data
		QuestionAnswer:   []structure.QuestionAnswer{},   // Populate with necessary data
		Bid:              []structure.Bid{},              // Populate with necessary data
		Review:           structure.Review{},           // Populate with necessary data
	}
	if err := Create(collectionName, &expectedJob1); err != nil {
		t.Fatalf("Failed to insert test job document 1: %v", err)
	}
	if err := Create(collectionName, &expectedJob2); err != nil {
		t.Fatalf("Failed to insert test job document 2: %v", err)
	}

	// Act
	var resultJobs []structure.Job
	err := Find(collectionName, bson.M{}, &resultJobs)

	// Assert
	if err != nil {
		t.Fatalf("Failed to retrieve job documents: %v", err)
	}
	if len(resultJobs) != 2 {
		t.Errorf("Expected 2 job documents, got %d", len(resultJobs))
	}
}
