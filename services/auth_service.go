// auth_service.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var authRequest AuthRequest
	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		logrus.Error("Error decoding auth request:", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	user, err := model.GetUser(authRequest.Login)
	if err != nil {
		logrus.Error("Error getting user:", err)
		http.Error(w, "Error getting user", http.StatusNotFound)
		return
	}

	if user.Password != authRequest.Password {
		logrus.Error("Incorrect password")
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	accessToken := GenerateAccessToken(user)
	response := AuthResponse{AccessToken: accessToken}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
