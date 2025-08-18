package voice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type VoiceConfig struct {
	Gender       texttospeechpb.SsmlVoiceGender
	VoiceName    string
	Rate         float64
	Pitch        float64
	VolumeGainDb float64
}

func SelectVoice() VoiceConfig {
	options := []string{
		"1. Arun (Male - Normal)",        	// Male (Normal)
		"2. Sindhu (Female - Normal)",    	// Female (Normal)
		"3. Pranav (Male - Announcement)",	// Male (Announcement)
		"4. Kavya (Female - Announcement)",	// Female (Announcement)
		"5. Harish (Male - Storytelling)",	// Male (Storytelling)
		"6. Manju (Female - Storytelling)",	// Female (Storytelling)
	}
	

	fmt.Println("Select a voice tone:")
	for _, opt := range options {
		fmt.Println(opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice (1-6): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, _ := strconv.Atoi(choiceStr)

	switch choice {
	case 1:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 1.0, 0.0, 0.0}
	case 2:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.0, 0.0, 0.0}
	case 3:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 1.2, 2.0, 3.0} // faster, louder
	case 4:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.2, 2.0, 3.0}
	case 5:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 0.9, -1.0, 0.0} // slower, deeper
	case 6:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 0.9, -1.0, 0.0}
	default:
		return VoiceConfig{texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.0, 0.0, 0.0}
	}
}
