package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(user, password, dbname, host string, port int, migrationsPath string) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", user, dbname, password, host, port)
	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error with db ping")
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}

// runMigrations applies database migrations
func RunMigrations(migrationsPath string) error {
	// Create a new migrate instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath), // Ensure migrationsPath is in the format "file://path/to/directory"
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not create migration instance: %w", err)
	}

	// Apply all up migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply up migrations: %w", err)
	}

	return nil
}
