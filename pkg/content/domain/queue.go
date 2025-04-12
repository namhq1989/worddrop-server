package domain

import "github.com/namhq1989/go-utilities/appcontext"

type QueueRepository interface {
	GenerateTextAudio(ctx *appcontext.AppContext, payload QueueGenerateTextAudioPayload) error
}

type QueueFetchNewsPayload struct {
	Service string
}

type QueueGenerateTextAudioPayload struct {
	ID      string
	Content string
}
