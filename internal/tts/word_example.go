package tts

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
)

func (t TTS) GenerateWordExampleAudio(ctx *appcontext.AppContext, exampleID, exampleContent string) (string, error) {
	var (
		fileName = fmt.Sprintf("%s.%s", exampleID, t.extension)
		voice    = t.randomVoice()
	)

	err := t.synthesizeAndUploadAudio(ctx, exampleContent, fileName, voice)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
