package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("./"))

	// Route for the static file server
	http.Handle("/", fs)

	// Start the server
	port := "8080"
	log.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

