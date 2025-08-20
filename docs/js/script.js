/**
 * ===============================
 * Dhwani Frontend JS – Malayalam TTS
 * ===============================
 * This file handles:
 * 1. User input for Malayalam text
 * 2. Voice selection
 * 3. Communication with the backend / Google TTS API
 * 4. Generating and playing audio
 * 5. Downloading the generated audio
 * 
 * Ready for frontend deployment with backend API on Cloud Run.
 */

document.addEventListener('DOMContentLoaded', () => {

    // -------------------------------
    // Element References
    // -------------------------------
    const audioOutput = document.getElementById('audio-output'); // Container for audio player & download
    const audioPlayer = document.getElementById('audio-player'); // <audio> element
    const downloadButton = document.getElementById('download-button'); // Download link
    const generateBtn = document.getElementById('generate-button'); // Generate Voice button
    const textInputMl = document.getElementById('text-input-ml'); // Malayalam text input
    const voiceSelect = document.getElementById('voice-select'); // Voice select dropdown
    const audioMessage = document.getElementById('audio-message'); // Status messages

    // -------------------------------
    // Voice Map – User selection to TTS voice
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
    // Helper: Convert Base64 to ArrayBuffer
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

    // -------------------------------
    // Helper: Convert PCM to WAV Blob
    // -------------------------------
    function pcmToWav(pcmData, sampleRate) {
        const buffer = new ArrayBuffer(44 + pcmData.byteLength);
        const view = new DataView(buffer);
        const littleEndian = true;

        function writeString(view, offset, string) {
            for (let i = 0; i < string.length; i++) {
                view.setUint8(offset + i, string.charCodeAt(i));
            }
        }

        writeString(view, 0, 'RIFF');                       // RIFF header
        view.setUint32(4, 36 + pcmData.byteLength, littleEndian); // File size
        writeString(view, 8, 'WAVE');                       // WAVE type
        writeString(view, 12, 'fmt ');                      // Format chunk
        view.setUint32(16, 16, littleEndian);              // Format chunk length
        view.setUint16(20, 1, littleEndian);               // Audio format = PCM
        view.setUint16(22, 1, littleEndian);               // Channels = 1
        view.setUint32(24, sampleRate, littleEndian);      // Sample rate
        view.setUint32(28, sampleRate * 2, littleEndian);  // Byte rate
        view.setUint16(32, 2, littleEndian);               // Block align
        view.setUint16(34, 16, littleEndian);              // Bits per sample
        writeString(view, 36, 'data');                     // Data chunk
        view.setUint32(40, pcmData.byteLength, littleEndian);

        const pcm16 = new Int16Array(pcmData);
        for (let i = 0; i < pcm16.length; i++) {
            view.setInt16(44 + i * 2, pcm16[i], littleEndian);
        }

        return new Blob([view], { type: 'audio/wav' });
    }

    // -------------------------------
    // Generate Voice Button Click
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

        // UI state: loading
        generateBtn.textContent = 'Generating...';
        generateBtn.disabled = true;
        audioOutput.classList.add('hidden');
        audioMessage.textContent = "";

        try {
            // -------------------------------
            // Backend API Call
            // -------------------------------
            // Replace this URL with your Cloud Run API if needed:
            const apiUrl = `/api/tts`;  // For Cloud Run deployment serving both frontend + backend

            const payload = {
                text: textToSpeak,
                voice: voiceName
            };

            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                throw new Error(`API call failed with status: ${response.status}`);
            }

            const result = await response.json();
            const audioBase64 = result.audioData; // Backend should return Base64 audio

            if (!audioBase64) throw new Error("No audio data received from API");

            // Convert Base64 → WAV → Play & Download
            const pcmData = base64ToArrayBuffer(audioBase64);
            const wavBlob = pcmToWav(pcmData, 24000); // Use sample rate from backend if dynamic
            const audioUrl = URL.createObjectURL(wavBlob);

            audioPlayer.src = audioUrl;
            downloadButton.href = audioUrl;

            audioMessage.textContent = "Your Malayalam voice is ready!";
            audioOutput.classList.remove('hidden');
            audioOutput.style.transform = 'scale(1)';
            audioOutput.style.opacity = '1';

        } catch (error) {
            console.error("Failed to generate voice:", error);
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
    // Custom Voice Dropdown Logic
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
            voiceSelect.value = selectedVoiceId; // Update hidden select input for API

            selectOptions.classList.remove('open');
            selectHeader.classList.remove('active');
        }
    });

    // Close dropdown if clicked outside
    document.addEventListener('click', (e) => {
        if (!selectHeader.contains(e.target) && !selectOptions.contains(e.target)) {
            selectOptions.classList.remove('open');
            selectHeader.classList.remove('active');
        }
    });

});
