package tts

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

func (t TTS) randomVoice() types.Voice {
	rand := manipulation.RandomIntInRange(0, len(t.voices)-1)
	return t.voices[rand]
}

func (t TTS) initDirectories() {
	if err := os.MkdirAll(t.getFilePath(), 0755); err != nil {
		panic(fmt.Errorf("failed to create TTS files directory %s: %s", t.getFilePath(), err.Error()))
	}
}

func (TTS) getFilePath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "files/tts")
}

func (t TTS) generateFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", t.getFilePath(), fileName)
}

func (t TTS) synthesizeAndUploadAudio(ctx *appcontext.AppContext, text string, fileName string, voice types.Voice) error {
	// Generate speech with Polly
	output, err := t.polly.SynthesizeSpeech(ctx.Context(), &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatMp3,
		Text:         aws.String(text),
		VoiceId:      voice.Id,
		Engine:       types.EngineNeural,
		LanguageCode: types.LanguageCodeEnUs,
		TextType:     types.TextTypeText,
	})
	if err != nil {
		ctx.Logger().Error("[tts] failed to synthesize speech from Polly", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = output.AudioStream.Close() }()

	// Create local file
	file, err := os.Create(t.generateFilePath(fileName))
	if err != nil {
		ctx.Logger().Error("[tts] failed to create file from Polly response", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = file.Close() }()

	// Write to local file
	_, err = io.Copy(file, output.AudioStream)
	if err != nil {
		ctx.Logger().Error("[tts] failed to write file from Polly response", err, appcontext.Fields{})
		return err
	}

	// Upload to R2
	err = t.uploadToR2(ctx, fileName)
	if err != nil {
		ctx.Logger().Error("[tts] failed to upload file to R2", err, appcontext.Fields{})
		return err
	}

	return nil
}

func (t TTS) uploadToR2(ctx *appcontext.AppContext, fileName string) error {
	localFilePath := t.generateFilePath(fileName)

	fileToUpload, err := os.Open(localFilePath)
	if err != nil {
		ctx.Logger().Error("[tts] failed to open file for R2 upload", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = fileToUpload.Close() }()

	_, err = t.r2.PutObject(ctx.Context(), &s3.PutObjectInput{
		Bucket:      aws.String(t.r2Bucket),
		Key:         aws.String(fileName),
		Body:        fileToUpload,
		ContentType: aws.String("audio/mp3"),
	})
	if err != nil {
		ctx.Logger().Error("[tts] failed to upload file to R2", err, appcontext.Fields{})
		return err
	}

	// remove local file
	err = os.Remove(localFilePath)
	if err != nil {
		ctx.Logger().Error("[tts] failed to remove local file", err, appcontext.Fields{})
		return err
	}

	return nil
}
