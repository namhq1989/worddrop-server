package tts

import (
	"fmt"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

func (t TTS) GenerateWordAudio(ctx *appcontext.AppContext, word string) (string, error) {
	var (
		slug     = manipulation.Slugify(word)
		fileName = fmt.Sprintf("%s.%s", slug, t.extension)
		voice    = t.randomVoice()
	)

	err := t.synthesizeAndUploadAudio(ctx, word, fileName, voice)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
