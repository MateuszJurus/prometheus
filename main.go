package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateuszjurus/prometheus/internal/handlers"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the route for the homepage
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// Serve static files from the React build directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/front/build/static"))))

	// Serve the main index.html file
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("web/front/build")))

	// Start the server on specified port
	const port string = ":8080"
	http.Handle("/", router)

	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
