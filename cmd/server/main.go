package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateuszjurus/prometheus/config"
	handlers "github.com/mateuszjurus/prometheus/pkg/api/handler"
	db "github.com/mateuszjurus/prometheus/pkg/store"
	"github.com/rs/cors"
)

func main() {
	connectDB(config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort, config.DBMigrationsPath)
	initRouting()
}

func initRouting() {
	// Create a new router
	router := mux.NewRouter()

	// Define the route for the homepage
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/list", handlers.ListUsersHandler).Methods("GET")

	// Use CORS middleware
	corsHandler := cors.Default().Handler(router)

	// Serve static files from the React build directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/front/build/static"))))

	// Serve the main index.html file
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("web/front/build")))

	// Start the server on specified port
	const port string = ":8080"
	http.Handle("/", router)

	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, corsHandler)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func connectDB(dbuser, dbpass, dbname, dbhost string, dbport int, dbmigrationspath string) {
	db.InitDB(dbuser, dbpass, dbname, dbhost, dbport, dbmigrationspath)
	db.RunMigrations(dbmigrationspath)
}
