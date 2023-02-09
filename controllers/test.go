// controllers/test.go
package controllers

import (
	"encoding/json"
	"github.com/cleisommais/oauth-service-v1/models"
	"net/http"
)

// TestController is a controller function that writes the response to a /test request
func TestController(w http.ResponseWriter, r *http.Request) {
	response := &models.TestResponse{Test: 123}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
