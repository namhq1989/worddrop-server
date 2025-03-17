package tts

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/namhq1989/go-utilities/appcontext"
	"google.golang.org/api/option"
)

type googleTTS struct {
	client        *texttospeech.Client
	neural2Voices []string
	chirp3Voices  []string
}

func newGoogleTTS(ctx context.Context, cfg GoogleConfig) *googleTTS {
	client, err := texttospeech.NewClient(ctx, option.WithAPIKey(cfg.ApiKey))
	if err != nil {
		panic(fmt.Errorf("failed to create Google TTS client: %v", err))
	}

	fmt.Printf("⚡️ [tts]: google TTS connected \n")

	return &googleTTS{
		client: client,
		// Neural2 voices - use in first half of month
		neural2Voices: []string{
			"en-US-Neural2-A",
			"en-US-Neural2-C",
			"en-US-Neural2-D",
			"en-US-Neural2-E",
			"en-US-Neural2-F",
			"en-US-Neural2-G",
			"en-US-Neural2-H",
			"en-US-Neural2-I",
			"en-US-Neural2-J",
		},
		// Chirp3 voices - use in second half of month
		chirp3Voices: []string{
			"en-US-Chirp3-HD-Aoede",
			"en-US-Chirp3-HD-Charon",
			"en-US-Chirp3-HD-Fenrir",
			"en-US-Chirp3-HD-Kore",
			"en-US-Chirp3-HD-Leda",
			"en-US-Chirp3-HD-Orus",
			"en-US-Chirp3-HD-Puck",
			"en-US-Chirp3-HD-Zephyr",
		},
	}
}

func (g googleTTS) randomVoice() string {
	now := time.Now()
	dayOfMonth := now.Day()

	var voices []string
	if dayOfMonth <= 15 {
		voices = g.neural2Voices
	} else {
		voices = g.chirp3Voices
	}

	randIndex := rand.Intn(len(voices))
	return voices[randIndex]
}

func (g googleTTS) synthesize(ctx *appcontext.AppContext, t TTS, text string, fileName string) error {
	voice := g.randomVoice()

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			Name:         voice,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := g.client.SynthesizeSpeech(ctx.Context(), &req)
	if err != nil {
		ctx.Logger().Error("[tts-google] failed to synthesize speech", err, appcontext.Fields{})
		return err
	}

	// Create the output file
	file, err := os.Create(t.generateFilePath(fileName))
	if err != nil {
		ctx.Logger().Error("[tts-google] failed to create file", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = file.Close() }()

	// Write the audio content to the file
	_, err = file.Write(resp.AudioContent)
	if err != nil {
		ctx.Logger().Error("[tts-google] failed to write file", err, appcontext.Fields{})
		return err
	}

	// Upload to R2
	if err = t.uploadToR2(ctx, fileName); err != nil {
		ctx.Logger().Error("[tts-google] failed to upload file to R2", err, appcontext.Fields{})
		return err
	}

	return nil
}
