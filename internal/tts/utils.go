package tts

import (
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/namhq1989/go-utilities/appcontext"
)

func (t TTS) isServicePolly() bool {
	return t.service == ServicePolly
}

func (t TTS) isServiceGoogle() bool {
	return t.service == ServiceGoogle
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

	if err = os.Remove(localFilePath); err != nil {
		ctx.Logger().Error("[tts] failed to remove local file", err, appcontext.Fields{})
		return err
	}

	return nil
}
