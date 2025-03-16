package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/tts"
)

type TTSRepository struct {
	tts tts.Operations
}

func NewTTSRepository(tts tts.Operations) TTSRepository {
	return TTSRepository{
		tts: tts,
	}
}

func (r TTSRepository) GenerateTextAudio(ctx *appcontext.AppContext, id, content string) error {
	return r.tts.GenerateTextAudio(ctx, id, content)
}
