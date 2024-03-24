// handlers/user_handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mateuszjurus/prometheus/pkg/domain"
	"github.com/mateuszjurus/prometheus/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Printf("Error hashing password: %v", err)
		return
	}
	newUser.Password = string(hashedPassword)

	// Insert the new user into the database
	err = store.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Printf("Error creating user: %v", err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := store.ListUsers()
	if err != nil {
		http.Error(w, "Error listing users", http.StatusInternalServerError)
		log.Printf("Error listing users: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Error encoding users to JSON", http.StatusInternalServerError)
		log.Printf("Error encoding users to JSON: %v", err)
	}
}
