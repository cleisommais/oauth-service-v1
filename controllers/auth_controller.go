// auth_controller.go
package AuthController

import (
	"encoding/json"
	"net/http"

	"github.com/oauth-service-v1/auth-microservice/model"
	"github.com/oauth-service-v1/auth-microservice/service"
)

// Authenticate handles the /auth endpoint
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Get the login and password from the request parameters
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")

	// Get the user from the database
	user, err := model.GetUser(login)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Validate the password
	if user.Password != password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT access token
	accessToken, err := service.GenerateAccessToken(user.ID)
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	// Return the access token in the response
	response := map[string]string{"access_token": accessToken}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
