package domain

import "github.com/namhq1989/go-utilities/appcontext"

type TtsRepository interface {
	GenerateTextAudio(ctx *appcontext.AppContext, id, content string) error
}
