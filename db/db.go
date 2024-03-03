package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(user, password, dbname, host string, port int) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", user, dbname, password, host, port)
	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error with db ping")
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}
