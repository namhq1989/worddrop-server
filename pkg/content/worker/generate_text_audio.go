package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"go.opentelemetry.io/otel"
)

type GenerateTextAudioHandler struct {
	ttsRepository domain.TtsRepository
}

func NewGenerateTextAudioHandler(ttsRepository domain.TtsRepository) GenerateTextAudioHandler {
	return GenerateTextAudioHandler{
		ttsRepository: ttsRepository,
	}
}

func (h GenerateTextAudioHandler) GenerateTextAudio(ctx *appcontext.AppContext, payload domain.QueueGenerateTextAudioPayload) error {
	tracer := otel.Tracer("[tracer] generate text audio")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] generate text audio")
	ctx.SetContext(spanCtx)
	defer span.End()

	return h.ttsRepository.GenerateTextAudio(ctx, payload.ID, payload.Content)
}
