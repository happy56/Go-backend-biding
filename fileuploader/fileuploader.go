package fileuploader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// Define the maximum file size allowed (in bytes)
const maxFileSize = 10 << 20 // 10 MB

// Define the allowed file types
var allowedFileTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
	".txt":  true,
	// Add more file types as needed
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form with a maximum file size
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Println("Error Parsing Multipart Form:", err)
		http.Error(w, "Error Parsing Multipart Form", http.StatusBadRequest)
		return
	}

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File:", err)
		http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check file size
	if handler.Size > maxFileSize {
		fmt.Println("File size exceeds the maximum allowed size")
		http.Error(w, "File size exceeds the maximum allowed size", http.StatusBadRequest)
		return
	}

	// Check file type
	ext := filepath.Ext(handler.Filename)
	if !allowedFileTypes[ext] {
		fmt.Println("File type not allowed")
		http.Error(w, "File type not allowed", http.StatusBadRequest)
		return
	}

	// Generate a new filename with only alphanumeric characters while keeping the extension
	newFilename := generateAlphanumericFilename(handler.Filename)

	fmt.Println("newFilename", newFilename)

	// Set the upload directory name
	uploadDir := "uploadedfiles"

	// Check if the directory exists, if not, create it
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			http.Error(w, "Error creating directory", http.StatusInternalServerError)
			return
		}
	}

	// Create a file within our upload directory with the new filename
	newFilePath := filepath.Join(uploadDir, newFilename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file content to the newly created file
	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		http.Error(w, "Error copying file content", http.StatusInternalServerError)
		return
	}

	// Return that we have successfully uploaded our file
	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", newFilename)
}

// generateAlphanumericFilename generates a new filename with only alphanumeric characters
func generateAlphanumericFilename(filename string) string {
	// Extract the extension
	ext := filepath.Ext(filename)

	// Extract the filename without the extension
	name := filename[:len(filename)-len(ext)]

	// Define a regular expression to match non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		return ""
	}

	// Replace non-alphanumeric characters with an empty string in the name
	alphanumericName := reg.ReplaceAllString(name, "")

	// Return the filename with the dot (".") before the extension
	return alphanumericName + ext
}
