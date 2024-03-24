package domain

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserSvc interface {
	GetUser(id int) (*User, error)
	ListUser(category string) ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

type UserDB interface {
	GetUser(id int) (*User, error)
	ListUser(category string) ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

func GetUserByUserName(db *sql.DB, username string) (*User, error) {
	var user User
	query := `SELECT "ID", username, email, password, role FROM users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			// No result, return a nil user and no error (or custom NotFound error)
			fmt.Printf("Was unable to find user: %s\n", username)
			return nil, nil
		}
		// An actual error occurred
		return nil, fmt.Errorf("error fetching user by username: %w", err)
	}

	return &user, nil
}
