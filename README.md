Here’s a polished, drop‑in README.md for your project. Tweak anything that doesn’t match your current code.

```markdown
# DHWANI — Malayalam Text‑to‑Speech (TTS) 🔊🇮🇳

DHWANI converts Malayalam text into natural‑sounding MP3 audio using Google Cloud Text‑to‑Speech. It ships with an HTTP API, multiple voice styles, and configurable rate, pitch, and volume.

> ധ്വനി (Dhwani) = “sound/voice” in Malayalam

---

## Features
- 🗣️ Convert Malayalam text to MP3
- 🎙️ Multiple voices: Male/Female + styles (Normal, Announcement, Storytelling)
- 🎛️ Adjustable speech rate, pitch, and volume
- 🌐 HTTP endpoints to generate and stream/download audio
- 🧩 Modular, extensible design

---

## Tech Stack
- Language: Go
- TTS Provider: Google Cloud Text‑to‑Speech
- HTTP Router: Chi
- Output: MP3

---

## Folder Structure
```
DHWANI/
├─ internal/
│  ├─ handler/
│  │  └─ ttsHandler.go      # HTTP endpoints (/generate, /play)
│  └─ tts/
│     └─ tts.go             # Google Cloud TTS integration
├─ textsource/
│  └─ text.go               # Example/default Malayalam text source
├─ voice/
│  └─ voice.go              # Voice selection (currently via stdin)
├─ output/                  # Generated .mp3 files
├─ .env                     # Environment variables (ignored)
├─ .gitignore
├─ go.mod
├─ go.sum
├─ main.go                  # App entrypoint
└─ README.md
```

---

## Prerequisites
- Go 1.21+ installed
- A Google Cloud project with the “Text‑to‑Speech API” enabled
- A Service Account JSON key with access to Text‑to‑Speech

> Costs: Google Cloud TTS may incur charges. Monitor your usage.

---

## Setup

1) Clone and install dependencies
```bash
git clone <your-repo-url> dhwani
cd dhwani
go mod tidy
```

2) Create a Service Account and download the JSON key  
In Google Cloud Console:
- Enable “Text‑to‑Speech API”
- Create Service Account → Assign suitable role (e.g., Cloud Text‑to‑Speech User)
- Download key as JSON

3) Create .env
```
GOOGLE_APPLICATION_CREDENTIALS=/absolute/path/to/credentials.json
PORT=8080
```

4) Run the server
```bash
go run main.go
```
Server starts on http://localhost:8080 (or PORT from .env).

---

## API

### 1) Generate audio
- Method: POST  
- URL: `/generate`  
- Body (JSON):
```json
{
  "text": "നമസ്കാരം! നിങ്ങൾക്ക്‌ സുഖംതന്നെ?",
  "voice": "female-story",
  "rate": 1.0,
  "pitch": 0.0,
  "volumeGainDb": 0.0
}
```

Notes:
- voice is optional; if omitted, the server uses its default voice (current code may still read from stdin; see “Current limitation” below).
- rate, pitch, volumeGainDb are optional and may be ignored if not wired in your version.

- Successful response:
```json
{
  "file": "output/output_1690000000.mp3"
}
```

- Example with curl:
```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"text":"നമസ്കാരം! നിങ്ങൾക്ക്‌ സുഖംതന്നെ?","voice":"female-story"}'
```

### 2) Play or download audio
- Method: GET  
- URL: `/play?file=<filename>`
- Example:
```bash
curl -L "http://localhost:8080/play?file=output_1690000000.mp3" --output dhwani.mp3
```
Open in browser:
```
http://localhost:8080/play?file=output_1690000000.mp3
```

---

## Voices

Available keys (extend as you like):
- `male`
- `female`
- `male-announcement`
- `female-story`

Behind the scenes, these map to Google Cloud voice names and speaking styles. Add/update mappings in `voice/voice.go`.

---

## Configuration

Environment variables:
- GOOGLE_APPLICATION_CREDENTIALS: Absolute path to the Service Account JSON key
- PORT: HTTP server port (default: 8080)

Other app defaults (voice, rate, pitch, etc.) live in `voice/voice.go` and `internal/tts/tts.go`.

---

## How it works

- `internal/tts/tts.go`
  - Builds the TTS request (voice, audio config)
  - Calls Google Cloud TTS and writes MP3 to `output/`
- `internal/handler/ttsHandler.go`
  - Exposes `/generate` and `/play`
  - Validates input and returns the generated file path
- `voice/voice.go`
  - Maps human‑friendly voice keys (e.g., `female-story`) to provider voices
  - Currently reads from stdin when run interactively
- `textsource/text.go`
  - Provides a default Malayalam text string for demo/testing

---

## Current limitation

The codebase notes that voice selection is currently done via stdin. If your HTTP body `voice` isn’t yet connected, the server may ignore it and fallback to interactive selection or defaults. To fully support HTTP‑driven voice selection:
- Read `voice`, `rate`, `pitch`, `volumeGainDb` from the POST body in `ttsHandler.go`
- Pass those values into the TTS layer instead of stdin

---

## Roadmap
- Accept voice and audio options fully via HTTP (no stdin)
- Input validation (length, allowed voices)
- Use UUIDs for filenames
- Stream audio directly (avoid writing to disk)
- Graceful shutdown
- Unit tests for handlers and TTS service
- Dockerfile for containerized deploy
- Caching for repeated texts

---

## Troubleshooting

- “google: could not find default credentials”
  - Ensure `GOOGLE_APPLICATION_CREDENTIALS` points to a valid JSON key file
  - Use an absolute path and restart the server
- Permission or 403 errors
  - Confirm the Service Account has Text‑to‑Speech permissions
- Empty or corrupted MP3
  - Check the request body and logs
  - Verify the voice mapping and audio config
- File not found on `/play`
  - Use the exact `file` name returned by `/generate`
  - Ensure files exist in the `output/` directory

---

## Security
- Do not commit `.env` or credential files
- Restrict Service Account permissions to the minimum required
- Consider rate‑limiting or auth if exposed publicly

---

## Contributing
PRs are welcome! Please:
1) Open an issue describing the change
2) Write tests where possible
3) Follow Go idioms and keep the architecture modular

---

## License
MIT — see LICENSE for details.
```

Want me to tailor this to your exact request/response schemas (e.g., final JSON field names, default voice, or add a Docker section)?