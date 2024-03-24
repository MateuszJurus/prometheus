package store

import (
	"fmt"
	"log"

	"github.com/mateuszjurus/prometheus/pkg/domain"
)

func CreateUser(user domain.User) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	} else {
		fmt.Println("User added successfully")
	}
	return nil
}
