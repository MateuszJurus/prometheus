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
