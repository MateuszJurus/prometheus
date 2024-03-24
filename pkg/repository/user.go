package repository

import (
	"database/sql"
	"log"

	"github.com/mateuszjurus/prometheus/pkg/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user domain.User) error {
	_, err := repo.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	}
	log.Println("User added successfully")
	return nil
}
