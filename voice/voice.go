package voice

import (
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type VoiceConfig struct {
	Name         string `json:"name"`
	Gender       texttospeechpb.SsmlVoiceGender
	VoiceName    string
	Rate         float64
	Pitch        float64
	VolumeGainDb float64
}

func GetVoiceConfig(choice int) VoiceConfig {
	switch choice {
	case 1:
		return VoiceConfig{"Arun (Male - Normal)", texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 1.0, 0.0, 0.0}
	case 2:
		return VoiceConfig{"Sindhu (Female - Normal)", texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.0, 0.0, 0.0}
	case 3:
		return VoiceConfig{"Pranav (Male - Announcement)", texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 1.2, 2.0, 3.0}
	case 4:
		return VoiceConfig{"Kavya (Female - Announcement)", texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.2, 2.0, 3.0}
	case 5:
		return VoiceConfig{"Harish (Male - Storytelling)", texttospeechpb.SsmlVoiceGender_MALE, "ml-IN-Standard-B", 0.9, -1.0, 0.0}
	case 6:
		return VoiceConfig{"Manju (Female - Storytelling)", texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 0.9, -1.0, 0.0}
	default:
		return VoiceConfig{"Default (Female)", texttospeechpb.SsmlVoiceGender_FEMALE, "ml-IN-Standard-A", 1.0, 0.0, 0.0}
	}
}

// ✅ New: return all voices for frontend listing
func ListVoices() []map[string]interface{} {
	return []map[string]interface{}{
		{"id": 1, "name": "Arun (Male - Normal)"},
		{"id": 2, "name": "Sindhu (Female - Normal)"},
		{"id": 3, "name": "Pranav (Male - Announcement)"},
		{"id": 4, "name": "Kavya (Female - Announcement)"},
		{"id": 5, "name": "Harish (Male - Storytelling)"},
		{"id": 6, "name": "Manju (Female - Storytelling)"},
	}
}
