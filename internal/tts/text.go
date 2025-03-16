package tts

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
)

func (t TTS) GenerateTextAudio(ctx *appcontext.AppContext, id, content string) error {
	var (
		fileName = fmt.Sprintf("%s.%s", id, t.extension)
		voice    = t.randomVoice()
	)

	return t.synthesizeAndUploadAudio(ctx, content, fileName, voice)
}
