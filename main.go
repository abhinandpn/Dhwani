package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/abhinandpn/Dhwani/textsource" // import your package

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func main() {
	ctx := context.Background()

	credsPath := "/home/delta/Downloads/Dhwani-GTTS/dhwani-469106-63cddd3273b0.json"

	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	text := textsource.GetText() // Get text from another package

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ml-IN", // Malayalam
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		log.Fatalf("Failed to synthesize speech: %v", err)
	}

	err = ioutil.WriteFile("output.mp3", resp.AudioContent, 0644)
	if err != nil {
		log.Fatalf("Failed to save audio: %v", err)
	}

	fmt.Println("Audio content written to 'output.mp3'")
}
