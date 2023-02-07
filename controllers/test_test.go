// controllers/test_test.go
package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestController(t *testing.T) {
	// Create a test request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the TestController function with the request and response recorder
	TestController(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("TestController returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("TestController returned wrong Content-Type: got %v want %v", contentType, "application/json")
	}

	// Check the body of the response
	expected := `{"test":123}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("TestController returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
