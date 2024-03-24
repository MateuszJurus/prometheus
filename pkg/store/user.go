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

func ListUsers() ([]domain.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error fetching users from DB: %v", err)
		return nil, err
	} else {
		fmt.Println("Success fetching users from DB")
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	} else {
		return users, nil
	}
}
