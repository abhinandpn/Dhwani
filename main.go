package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	text := "നമസ്കാരം, ഇത് വളരെ വ്യക്തവും മനോഹരവുമായ സ്ത്രീ ശബ്ദത്തിൽ നിർമ്മിച്ച ടെക്സ്റ്റ് ടു സ്പീച്ച് ഓഡിയോ ആണ്."

	// Change model_name if you have a specific one from Python
	modelName := "tts_models/en/ljspeech/tacotron2-DDC_ph" // Replace with your female voice model
	outputFile := "output.wav"

	cmd := exec.Command(
		"tts",
		"--text", text,
		"--model_name", modelName,
		"--out_path", outputFile,
	)

	fmt.Println("🎤 Generating high-quality female voice...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Error running TTS: %v\nDetails: %s", err, string(output))
	}

	fmt.Printf("✅ Audio saved to %s\n", outputFile)
}
