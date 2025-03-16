package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/queue"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type QueueRepository struct {
	queue queue.Operations
}

func NewQueueRepository(queue queue.Operations) QueueRepository {
	return QueueRepository{
		queue: queue,
	}
}

func (r QueueRepository) GenerateTextAudio(ctx *appcontext.AppContext, payload domain.QueueGenerateTextAudioPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.GenerateTextAudio, payload, 3)
}
