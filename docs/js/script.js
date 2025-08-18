document.addEventListener('DOMContentLoaded', () => {
    const audioOutput = document.getElementById('audio-output');
    const audioPlayer = document.getElementById('audio-player');
    const downloadButton = document.getElementById('download-button');
    const generateBtn = document.getElementById('generate-button');
    const textInputMl = document.getElementById('text-input-ml');
    const voiceSelect = document.getElementById('voice-select');
    const audioMessage = document.getElementById('audio-message');
    
    // Map the selected voice option to a Gemini TTS voice name
    const voiceMap = {
        'voice1': 'Zephyr',
        'voice2': 'Puck',
        'voice3': 'Kore',
        'voice4': 'Leda',
        'voice5': 'Orus',
        'voice6': 'Sadachbia',
    };

    // Helper function to convert base64 to ArrayBuffer
    function base64ToArrayBuffer(base64) {
        const binaryString = window.atob(base64);
        const len = binaryString.length;
        const bytes = new Uint8Array(len);
        for (let i = 0; i < len; i++) {
            bytes[i] = binaryString.charCodeAt(i);
        }
        return bytes.buffer;
    }

    // Helper function to convert raw PCM audio to a WAV file Blob
    function pcmToWav(pcmData, sampleRate) {
        const buffer = new ArrayBuffer(44 + pcmData.byteLength);
        const view = new DataView(buffer);
        const littleEndian = true;

        /* RIFF identifier */
        writeString(view, 0, 'RIFF');
        /* RIFF chunk length */
        view.setUint32(4, 36 + pcmData.byteLength, littleEndian);
        /* RIFF type */
        writeString(view, 8, 'WAVE');
        /* format chunk identifier */
        writeString(view, 12, 'fmt ');
        /* format chunk length */
        view.setUint32(16, 16, littleEndian);
        /* sample format (raw) */
        view.setUint16(20, 1, littleEndian);
        /* channel count */
        view.setUint16(22, 1, littleEndian);
        /* sample rate */
        view.setUint32(24, sampleRate, littleEndian);
        /* byte rate (sample rate * block align) */
        view.setUint32(28, sampleRate * 2, littleEndian);
        /* block align (channel count * bytes per sample) */
        view.setUint16(32, 2, littleEndian);
        /* bits per sample */
        view.setUint16(34, 16, littleEndian);
        /* data chunk identifier */
        writeString(view, 36, 'data');
        /* data chunk length */
        view.setUint32(40, pcmData.byteLength, littleEndian);

        // Write the PCM data
        const pcm16 = new Int16Array(pcmData);
        for (let i = 0; i < pcm16.length; i++) {
            view.setInt16(44 + i * 2, pcm16[i], littleEndian);
        }

        return new Blob([view], { type: 'audio/wav' });

        function writeString(view, offset, string) {
            for (let i = 0; i < string.length; i++) {
                view.setUint8(offset + i, string.charCodeAt(i));
            }
        }
    }


    // Handle the "Generate Voice" button click
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
        
        // UI state: show loading
        generateBtn.textContent = 'Generating...';
        generateBtn.disabled = true;
        audioOutput.classList.add('hidden');
        audioOutput.style.transform = 'scale(0.95)';
        audioOutput.style.opacity = '0';
        audioMessage.textContent = "";

        try {
            const payload = {
                contents: [{
                    parts: [{ text: textToSpeak }]
                }],
                generationConfig: {
                    responseModalities: ["AUDIO"],
                    speechConfig: {
                        voiceConfig: {
                            prebuiltVoiceConfig: { voiceName: voiceName }
                        }
                    }
                },
                model: "gemini-2.5-flash-preview-tts"
            };

            const apiKey = "";
            const apiUrl = `https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash-preview-tts:generateContent?key=${apiKey}`;

            // Exponential backoff for retries
            let retries = 0;
            const maxRetries = 5;
            const baseDelay = 1000;

            let response;
            while (retries < maxRetries) {
                try {
                    response = await fetch(apiUrl, {
                        method: 'POST',
                        headers: { 'Content-Type': 'application/json' },
                        body: JSON.stringify(payload)
                    });
                    if (response.status === 429) { // Too many requests
                        const delay = baseDelay * Math.pow(2, retries);
                        await new Promise(res => setTimeout(res, delay));
                        retries++;
                        continue;
                    }
                    break; // Success or non-retryable error
                } catch (e) {
                    throw e;
                }
            }

            if (!response || !response.ok) {
                throw new Error(`API call failed with status: ${response ? response.status : 'No response'}`);
            }

            const result = await response.json();
            const part = result?.candidates?.[0]?.content?.parts?.[0];
            const audioData = part?.inlineData?.data;
            const mimeType = part?.inlineData?.mimeType;

            if (audioData && mimeType && mimeType.startsWith("audio/L16")) {
                const sampleRate = parseInt(mimeType.match(/rate=(\d+)/)[1], 10);
                const pcmData = base64ToArrayBuffer(audioData);
                const wavBlob = pcmToWav(pcmData, sampleRate);
                const audioUrl = URL.createObjectURL(wavBlob);
                
                // Update UI with generated audio
                audioPlayer.src = audioUrl;
                downloadButton.href = audioUrl;
                
                audioMessage.textContent = "Your Malayalam voice is ready!";
                audioOutput.classList.remove('hidden');
                audioOutput.style.transform = 'scale(1)';
                audioOutput.style.opacity = '1';
            } else {
                 throw new Error("Invalid audio data received from API.");
            }

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
});

document.addEventListener('DOMContentLoaded', () => {
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
        
        // Here you can use the selectedVoiceId to perform actions
        console.log(`Voice selected: ${selectedText} (ID: ${selectedVoiceId})`);
        
        selectOptions.classList.remove('open');
        selectHeader.classList.remove('active');
      }
    });
  
    // Close the dropdown if the user clicks outside of it
    document.addEventListener('click', (e) => {
      if (!selectHeader.contains(e.target) && !selectOptions.contains(e.target)) {
        selectOptions.classList.remove('open');
        selectHeader.classList.remove('active');
      }
    });
  });