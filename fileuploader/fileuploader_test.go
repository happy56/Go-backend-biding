package fileuploader

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
	// Create a sample file to upload
	requestBody := strings.NewReader("--boundary\r\nContent-Disposition: form-data; name=\"myFile\"; filename=\"test.txt\"\r\nContent-Type: text/plain\r\n\r\nTest file content\r\n--boundary--")
	req := httptest.NewRequest("POST", "http://localhost:8080/upload", requestBody)
	req.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")
	res := httptest.NewRecorder()

	// Call the uploadFile function directly, passing the recorder and request.
	UploadFile(res, req)

	// Check the status code of the response.
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// Check the response body.
	expectedResponseBody := "Successfully Uploaded File: test.txt\n"
	if res.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body %q but got %q", expectedResponseBody, res.Body.String())
	}
}
