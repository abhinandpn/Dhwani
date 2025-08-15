package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/abhinandpn/Dhwani/textsource"
	"github.com/abhinandpn/Dhwani/voice"

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

	// Get Malayalam text from package
	text := textprovider.GetText()
	fmt.Println("Malayalam Text:", text)

	// Get voice from package
	gender := voice.SelectVoice()

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ml-IN", // Malayalam language code
			SsmlGender:   gender,
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

	fmt.Println("✅ Audio content written to 'output.mp3'")
}
