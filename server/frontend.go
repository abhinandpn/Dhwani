package server

import (
	"log"
	"net/http"
)

// Frontend serves your static files from the "docs" folder
func Frontend() error {
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/", fs)

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
	return nil
}
