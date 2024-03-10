// db/user_db.go

package db

import (
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user User) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	} else {
		fmt.Println("User added successfully")
	}
	return nil
}
