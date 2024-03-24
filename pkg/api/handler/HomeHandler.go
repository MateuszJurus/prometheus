package handlers

import (
	"net/http"
	"path/filepath"
)

// Home is the handler for the home route "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path for the React build directory
	reactBuildPath := "web/build"

	// Serve the index.html file
	http.ServeFile(w, r, filepath.Join(reactBuildPath, "index.html"))
}
