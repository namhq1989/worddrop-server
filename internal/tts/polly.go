package tts

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

var pollyPickedVoices = []string{"Joey", "Kendra", "Salli", "Ruth", "Stephen", "Gregory"}

type pollyTTS struct {
	client *polly.Client
	voices []types.Voice
}

func newPolly(ctx context.Context, awsCfg AwsConfig) *pollyTTS {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(awsCfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsCfg.AccessKey, awsCfg.SecretKey, "")),
	)
	if err != nil {
		panic(err)
	}
	svc := polly.NewFromConfig(cfg)

	fmt.Printf("⚡️ [tts]: polly TTS connected \n")

	input := &polly.DescribeVoicesInput{LanguageCode: types.LanguageCodeEnUs, Engine: types.EngineNeural}
	resp, err := svc.DescribeVoices(ctx, input)
	if err != nil {
		panic(err)
	}
	if len(resp.Voices) == 0 {
		panic(errors.New("no available voices of polly"))
	}

	voices := make([]types.Voice, 0)

	for _, voice := range resp.Voices {
		index := slices.IndexFunc(pollyPickedVoices, func(v string) bool {
			return v == *voice.Name
		})
		if index != -1 {
			voices = append(voices, voice)
		}
	}

	return &pollyTTS{
		client: svc,
		voices: voices,
	}
}

func (p pollyTTS) randomVoice() types.Voice {
	rand := manipulation.RandomIntInRange(0, len(p.voices)-1)
	return p.voices[rand]
}

func (p pollyTTS) synthesize(ctx *appcontext.AppContext, t TTS, text string, fileName string) error {
	voice := p.randomVoice()

	output, err := p.client.SynthesizeSpeech(ctx.Context(), &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatMp3,
		Text:         aws.String(text),
		VoiceId:      voice.Id,
		Engine:       types.EngineNeural,
		LanguageCode: types.LanguageCodeEnUs,
		TextType:     types.TextTypeText,
	})
	if err != nil {
		ctx.Logger().Error("[tts-polly] failed to synthesize speech from Polly", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = output.AudioStream.Close() }()

	file, err := os.Create(t.generateFilePath(fileName))
	if err != nil {
		ctx.Logger().Error("[tts-polly] failed to create file from Polly response", err, appcontext.Fields{})
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, output.AudioStream)
	if err != nil {
		ctx.Logger().Error("[tts-polly] failed to write file from Polly response", err, appcontext.Fields{})
		return err
	}

	// Upload to R2
	if err = t.uploadToR2(ctx, fileName); err != nil {
		ctx.Logger().Error("[tts-polly] failed to upload file to R2", err, appcontext.Fields{})
		return err
	}

	return nil
}
