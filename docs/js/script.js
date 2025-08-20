/**
 * ===============================
 * Dhwani Frontend JS – Malayalam TTS
 * ===============================
 * Handles:
 * 1. User input (Malayalam text)
 * 2. Voice selection
 * 3. Backend API call (Cloud Run)
 * 4. Generate + play audio dynamically (uses backend sampleRate)
 * 5. Download audio
 */

document.addEventListener('DOMContentLoaded', () => {

    // -------------------------------
    // Element References
    // -------------------------------
    const audioOutput = document.getElementById('audio-output');
    const audioPlayer = document.getElementById('audio-player');
    const downloadButton = document.getElementById('download-button');
    const generateBtn = document.getElementById('generate-button');
    const textInputMl = document.getElementById('text-input-ml');
    const voiceSelect = document.getElementById('voice-select');
    const audioMessage = document.getElementById('audio-message');

    // -------------------------------
    // Voice Map – user selection → TTS voice
    // -------------------------------
    const voiceMap = {
        'voice1': 'Zephyr',
        'voice2': 'Puck',
        'voice3': 'Kore',
        'voice4': 'Leda',
        'voice5': 'Orus',
        'voice6': 'Sadachbia',
    };

    // -------------------------------
    // Helpers
    // -------------------------------
    function base64ToArrayBuffer(base64) {
        const binaryString = window.atob(base64);
        const len = binaryString.length;
        const bytes = new Uint8Array(len);
        for (let i = 0; i < len; i++) {
            bytes[i] = binaryString.charCodeAt(i);
        }
        return bytes.buffer;
    }

    function pcmToWav(pcmData, sampleRate) {
        const buffer = new ArrayBuffer(44 + pcmData.byteLength);
        const view = new DataView(buffer);
        const littleEndian = true;

        function writeString(view, offset, string) {
            for (let i = 0; i < string.length; i++) {
                view.setUint8(offset + i, string.charCodeAt(i));
            }
        }

        writeString(view, 0, 'RIFF');
        view.setUint32(4, 36 + pcmData.byteLength, littleEndian);
        writeString(view, 8, 'WAVE');
        writeString(view, 12, 'fmt ');
        view.setUint32(16, 16, littleEndian);
        view.setUint16(20, 1, littleEndian);  // PCM
        view.setUint16(22, 1, littleEndian);  // Mono
        view.setUint32(24, sampleRate, littleEndian);
        view.setUint32(28, sampleRate * 2, littleEndian);
        view.setUint16(32, 2, littleEndian);
        view.setUint16(34, 16, littleEndian);
        writeString(view, 36, 'data');
        view.setUint32(40, pcmData.byteLength, littleEndian);

        const pcm16 = new Int16Array(pcmData);
        for (let i = 0; i < pcm16.length; i++) {
            view.setInt16(44 + i * 2, pcm16[i], littleEndian);
        }

        return new Blob([view], { type: 'audio/wav' });
    }

    // -------------------------------
    // Generate Voice
    // -------------------------------
    generateBtn.addEventListener('click', async () => {
        const textToSpeak = textInputMl.value.trim();
        if (!textToSpeak) {
            audioMessage.textContent = "Please enter some Malayalam text to generate audio.";
            audioOutput.classList.remove('hidden');
            audioOutput.style.transform = 'scale(1)';
            audioOutput.style.opacity = '1';
            return;
        }

        const selectedVoice = voiceSelect.value;
        const voiceName = voiceMap[selectedVoice];

        // Loading state
        generateBtn.textContent = 'Generating...';
        generateBtn.disabled = true;
        audioOutput.classList.add('hidden');
        audioMessage.textContent = "";

        try {
            // Call backend (Cloud Run serves both frontend + backend)
            const apiUrl = `/api/tts`;

            const payload = { text: textToSpeak, voice: voiceName };

            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            if (!response.ok) throw new Error(`API failed with status: ${response.status}`);

            const result = await response.json();
            const audioBase64 = result.audioData;
            const sampleRate = result.sampleRate || 24000; // fallback if backend doesn't return it

            if (!audioBase64) throw new Error("No audio data received from API");

            // Convert to playable WAV
            const pcmData = base64ToArrayBuffer(audioBase64);
            const wavBlob = pcmToWav(pcmData, sampleRate);
            const audioUrl = URL.createObjectURL(wavBlob);

            // Attach to player + download
            audioPlayer.src = audioUrl;
            downloadButton.href = audioUrl;
            downloadButton.download = "dhwani_audio.wav";

            // Show success
            audioMessage.textContent = "Your Malayalam voice is ready!";
            audioOutput.classList.remove('hidden');
            audioOutput.style.transform = 'scale(1)';
            audioOutput.style.opacity = '1';

        } catch (error) {
            console.error("TTS generation failed:", error);
            audioMessage.textContent = `Error: ${error.message}. Please try again.`;
            audioOutput.classList.remove('hidden');
            audioOutput.style.transform = 'scale(1)';
            audioOutput.style.opacity = '1';
        } finally {
            generateBtn.textContent = 'Generate Voice ✨';
            generateBtn.disabled = false;
        }
    });

    // -------------------------------
    // Custom Voice Dropdown
    // -------------------------------
    const selectHeader = document.querySelector('.select-header');
    const selectOptions = document.querySelector('.select-options');
    const selectedVoiceSpan = document.querySelector('.selected-voice');

    selectHeader.addEventListener('click', () => {
        selectOptions.classList.toggle('open');
        selectHeader.classList.toggle('active');
    });

    selectOptions.addEventListener('click', (e) => {
        if (e.target.tagName === 'LI') {
            const selectedText = e.target.textContent;
            const selectedVoiceId = e.target.dataset.voiceId;

            selectedVoiceSpan.textContent = selectedText;
            voiceSelect.value = selectedVoiceId;

            selectOptions.classList.remove('open');
            selectHeader.classList.remove('active');
        }
    });

    document.addEventListener('click', (e) => {
        if (!selectHeader.contains(e.target) && !selectOptions.contains(e.target)) {
            selectOptions.classList.remove('open');
            selectHeader.classList.remove('active');
        }
    });

});
