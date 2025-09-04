package app

import (
	"log"
	"net/http"
	"os"

	"github.com/abhinandpn/Dhwani/internal/handler"
	"github.com/abhinandpn/Dhwani/internal/tts"
	"github.com/abhinandpn/Dhwani/server"
	"github.com/go-chi/chi/v5"
)

// App holds the router and TTS service
type App struct {
	Router *chi.Mux
	Port   string
	Service *tts.TTSService
}

// NewApp initializes TTS service, router, and returns an App
func NewApp() (*App, error) {
	port := getPort()
	credsPath := getCredentialsPath()

	// Initialize TTS service
	service, err := tts.NewTTSService(credsPath)
	if err != nil {
		return nil, err
	}

	// Setup router with handlers
	router := setupRouter(handler.NewTTSHandler(service))

	return &App{
		Router: router,
		Port:   port,
		Service: service,
	}, nil
}

// Start launches the HTTP server
func (a *App) Start() {
	log.Printf("Server running on port %s", a.Port)
	if err := http.ListenAndServe(":"+a.Port, a.Router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// -------------------- Helpers --------------------

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func getCredentialsPath() string {
	credsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credsPath == "" {
		credsPath = "/Users/abhinanpn/Desktop/Credentials/dhwani-469106-63cddd3273b0.json"
	}
	return credsPath
}

func setupRouter(ttsHandler *handler.TTSHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Handle("/*", server.FrontendHandler())
	r.Post("/api/tts", ttsHandler.GenerateTTS)
	r.Get("/api/play", ttsHandler.PlayAudio)
	return r
}
