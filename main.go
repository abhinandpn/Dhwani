package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abhinandpn/Dhwani/internal/handler"
	"github.com/abhinandpn/Dhwani/internal/tts"
	"github.com/abhinandpn/Dhwani/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	// -------------------------------
	// Cloud Run dynamic PORT
	// -------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// -------------------------------
	// Google Credentials Path
	// -------------------------------
	// First, try to read from ENV (works inside Docker)
	credsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credsPath == "" {
		// fallback to your local dev path
		credsPath = "/home/delta/Downloads/Dhwani-GTTS/dhwani-469106-63cddd3273b0.json"
	}

	// -------------------------------
	// Initialize TTS service
	// -------------------------------
	service, err := tts.NewTTSService(credsPath)
	if err != nil {
		log.Fatalf("Failed to create TTS service (credsPath=%s): %v", credsPath, err)
	}

	// -------------------------------
	// Initialize TTS handler
	// -------------------------------
	ttsHandler := handler.NewTTSHandler(service)

	// -------------------------------
	// Router setup using Chi
	// -------------------------------
	r := chi.NewRouter()

	// Serve frontend files at "/"
	r.Handle("/*", server.FrontendHandler())

	// API endpoints
	r.Post("/api/tts", ttsHandler.GenerateTTS)
	r.Get("/api/play", ttsHandler.PlayAudio)

	// -------------------------------
	// Start server
	// -------------------------------
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
