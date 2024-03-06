// handlers/user_handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mateuszjurus/prometheus/db"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser db.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	fmt.Println(newUser)
	err = db.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}
