package tts

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
)

func (t TTS) GenerateTextAudio(ctx *appcontext.AppContext, id, content string) error {
	fileName := fmt.Sprintf("%s.%s", id, t.extension)

	if t.isServiceGoogle() {
		return t.google.synthesize(ctx, t, content, fileName)
	} else {
		return t.polly.synthesize(ctx, t, content, fileName)
	}
}
