package voice

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// SelectVoice lets the user choose a voice gender
func SelectVoice() texttospeechpb.SsmlVoiceGender {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select voice (male/female): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.ToLower(strings.TrimSpace(choice))

	if choice == "male" {
		return texttospeechpb.SsmlVoiceGender_MALE
	}
	return texttospeechpb.SsmlVoiceGender_FEMALE
}
