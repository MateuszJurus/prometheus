package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mateuszjurus/prometheus/pkg/auth"
	"github.com/mateuszjurus/prometheus/pkg/domain"
	db "github.com/mateuszjurus/prometheus/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	connection := db.GetDB()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := domain.GetUserByUserName(connection, req.Username)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Compare the password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Password does not match
		fmt.Println("Invalid credentials")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		fmt.Println("Failed to generate token")
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
