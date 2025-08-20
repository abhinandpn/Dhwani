package handler

import (
	"encoding/json"
	"net/http"

	"github.com/abhinandpn/Dhwani/internal/tts"

	"github.com/abhinandpn/Dhwani/internal/voice"
)

type TTSHandler struct {
	Service *tts.TTSService
}

func NewTTSHandler(service *tts.TTSService) *TTSHandler {
	return &TTSHandler{Service: service}
}

func (h *TTSHandler) GenerateTTS(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Text  string `json:"text"`
		Voice int    `json:"voice"` // expecting 1-6
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg := voice.GetVoiceConfig(body.Voice)

	filePath, err := h.Service.Synthesize(body.Text, cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"file": filePath}
	json.NewEncoder(w).Encode(resp)
}

func (h *TTSHandler) PlayAudio(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, filePath)
}
