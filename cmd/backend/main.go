package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler is the handler for the home route "/"
func Home(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	fmt.Fprint(w, "Hello World")
}

func main() {
	// Create new router using gorilla/mux
	router := mux.NewRouter()

	// Define a route and handler
	router.HandleFunc("/", Home).Methods("GET")

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
