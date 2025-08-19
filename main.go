package main

import (
	"log"
	"net/http"

	"github.com/abhinandpn/Dhwani/internal/handler"
	"github.com/abhinandpn/Dhwani/internal/tts"
	"github.com/go-chi/chi/v5"
)

func main() {

	// server.Frontend()
	// Path to your Google TTS JSON key
	credsPath := "/home/delta/Downloads/Dhwani-GTTS/dhwani-469106-63cddd3273b0.json"

	// Create TTS service
	service, err := tts.NewTTSService(credsPath)
	if err != nil {
		log.Fatalf("Failed to create TTS service: %v", err)
	}

	// Create handler
	ttsHandler := handler.NewTTSHandler(service)

	// Router setup
	r := chi.NewRouter()
	r.Post("/generate", ttsHandler.GenerateTTS)
	r.Get("/play", ttsHandler.PlayAudio)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
