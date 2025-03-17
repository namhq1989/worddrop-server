package tts

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	GenerateTextAudio(ctx *appcontext.AppContext, id, content string) error
}

const (
	ServicePolly  = "polly"
	ServiceGoogle = "google"
)

type TTS struct {
	extension string
	service   string
	polly     *pollyTTS
	google    *googleTTS
	r2        *s3.Client
	r2Bucket  string
}

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region    string
}

type GoogleConfig struct {
	ApiKey string
}

type R2Config struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
}

func NewTTSClient(usedService string, awsCfg AwsConfig, googleCfg GoogleConfig, r2Cfg R2Config) *TTS {
	var (
		ctx = context.Background()
		t   = &TTS{
			extension: "mp3",
			service:   usedService,
		}
	)

	if t.isServicePolly() {
		t.polly = newPolly(ctx, awsCfg)
	} else if t.isServiceGoogle() {
		t.google = newGoogleTTS(ctx, googleCfg)
	} else {
		panic("unknown TTS service")
	}

	initR2(ctx, t, r2Cfg)
	t.initDirectories()

	return t
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
