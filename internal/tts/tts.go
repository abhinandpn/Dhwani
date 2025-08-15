package tts

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type TTSService struct {
	Client *texttospeech.Client
}

func NewTTSService(credsPath string) (*TTSService, error) {
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		return nil, err
	}
	return &TTSService{Client: client}, nil
}

func (t *TTSService) Synthesize(text string, voiceType string) (string, error) {
	ctx := context.Background()

	// Voice selection based on request
	voiceParams := &texttospeechpb.VoiceSelectionParams{
		LanguageCode: "ml-IN",
		SsmlGender:   getGender(voiceType),
	}

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice:       voiceParams,
		AudioConfig: &texttospeechpb.AudioConfig{AudioEncoding: texttospeechpb.AudioEncoding_MP3},
	}

	resp, err := t.Client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return "", err
	}

	// Save file
	fileName := fmt.Sprintf("output/output_%d.mp3", time.Now().Unix())
	err = ioutil.WriteFile(fileName, resp.AudioContent, 0644)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func getGender(voiceType string) texttospeechpb.SsmlVoiceGender {
	switch voiceType {
	case "male", "male-announcement", "male-story":
		return texttospeechpb.SsmlVoiceGender_MALE
	default:
		return texttospeechpb.SsmlVoiceGender_FEMALE
	}
}
