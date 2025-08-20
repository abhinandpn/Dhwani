package tts

import (
	"context"
	"fmt"
	"os"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/abhinandpn/Dhwani/internal/voice"
	"google.golang.org/api/option"
)

type TTSService struct {
	Client *texttospeech.Client
}

// Existing function (manual path)
func NewTTSService(credsPath string) (*TTSService, error) {
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		return nil, err
	}
	return &TTSService{Client: client}, nil
}

// 🔹 New function (auto from env)
func NewTTSServiceFromEnv() (*TTSService, error) {
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credPath == "" {
		return nil, fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS not set")
	}
	return NewTTSService(credPath)
}

// Synthesize takes text + VoiceConfig and generates audio
func (t *TTSService) Synthesize(text string, cfg voice.VoiceConfig) (string, error) {
	ctx := context.Background()

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ml-IN",
			Name:         cfg.VoiceName,
			SsmlGender:   cfg.Gender,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			SpeakingRate:  cfg.Rate,
			Pitch:         cfg.Pitch,
			VolumeGainDb:  cfg.VolumeGainDb,
		},
	}

	resp, err := t.Client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return "", err
	}

	// Ensure output directory exists
	_ = os.MkdirAll("output", os.ModePerm)

	// Save file with timestamp
	fileName := fmt.Sprintf("output/output_%d.mp3", time.Now().Unix())
	err = os.WriteFile(fileName, resp.AudioContent, 0644)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
