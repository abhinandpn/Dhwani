package server

import (
	"net/http"
)

// FrontendHandler returns an http.Handler to serve static frontend files
// Serves all files from the "docs" folder
func FrontendHandler() http.Handler {
	fs := http.FileServer(http.Dir("./docs"))
	return fs
}
