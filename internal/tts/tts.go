package tts

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	GenerateWordAudio(ctx *appcontext.AppContext, word string) (string, error)
	GenerateWordExampleAudio(ctx *appcontext.AppContext, exampleID, exampleContent string) (string, error)
}

type TTS struct {
	extension string
	polly     *polly.Client
	voices    []types.Voice
	r2        *s3.Client
	r2Bucket  string
}

var pickedVoices = []string{"Joey", "Kendra", "Salli", "Ruth", "Stephen", "Gregory"}

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region    string
}

type R2Config struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
}

func NewTTSClient(awsCfg AwsConfig, r2Cfg R2Config) *TTS {
	var (
		ctx = context.Background()
		t   = &TTS{
			extension: "mp3",
		}
	)

	initPolly(ctx, t, awsCfg)
	initR2(ctx, t, r2Cfg)

	t.initDirectories()

	return t
}

func initPolly(ctx context.Context, t *TTS, awsCfg AwsConfig) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(awsCfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsCfg.AccessKey, awsCfg.SecretKey, "")),
	)
	if err != nil {
		panic(err)
	}
	svc := polly.NewFromConfig(cfg)

	fmt.Printf("⚡️ [tts]: polly connected \n")

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
		index := slices.IndexFunc(pickedVoices, func(v string) bool {
			return v == *voice.Name
		})
		if index != -1 {
			voices = append(voices, voice)
		}
	}

	t.polly = svc
	t.voices = voices
}

func initR2(ctx context.Context, t *TTS, r2Cfg R2Config) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2Cfg.AccessKey, r2Cfg.SecretKey, "")),
	)
	if err != nil {
		panic(err)
	}

	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(r2Cfg.Endpoint)
	})

	fmt.Printf("⚡️ [tts]: r2 connected \n")

	t.r2 = svc
	t.r2Bucket = r2Cfg.Bucket
}
