package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/queue"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type (
	Handlers interface {
		GenerateTextAudio(ctx *appcontext.AppContext, payload domain.QueueGenerateTextAudioPayload) error
	}
	Cronjob interface {
		FetchNews(ctx *appcontext.AppContext, _ domain.QueueFetchNewsPayload) error
	}
	Instance interface {
		Handlers
		Cronjob
	}

	workerHandlers struct {
		GenerateTextAudioHandler
	}
	workerCronjob struct {
		FetchNewsHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
		workerCronjob
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	wordRepository domain.WordRepository,
	wordNewsRepository domain.WordNewsRepository,
	wordExampleRepository domain.WordExampleRepository,
	externalAPIRepository domain.ExternalAPIRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
	ttsRepository domain.TtsRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			GenerateTextAudioHandler: NewGenerateTextAudioHandler(
				ttsRepository,
			),
		},
		workerCronjob: workerCronjob{
			FetchNewsHandler: NewFetchNewsHandler(
				wordRepository,
				wordNewsRepository,
				wordExampleRepository,
				externalAPIRepository,
				nlpRepository,
				queueRepository,
			),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GenerateTextAudio), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueGenerateTextAudioPayload](bgCtx, t, queue.ParsePayload[domain.QueueGenerateTextAudioPayload], w.GenerateTextAudio)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.FetchNews), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueFetchNewsPayload](bgCtx, t, queue.ParsePayload[domain.QueueFetchNewsPayload], w.FetchNews)
	})
}

type cronjobData struct {
	Task       string      `json:"task"`
	CronSpec   string      `json:"cronSpec"`
	Payload    interface{} `json:"payload"`
	RetryTimes int         `json:"retryTimes"`
}

func (w Worker) addCronjob() {
	var (
		ctx  = appcontext.NewWorker(context.Background())
		jobs = []cronjobData{
			{
				Task:       w.queue.GenerateTypename(queue.TypeNames.FetchNews),
				CronSpec:   "@every 2h",
				Payload:    domain.QueueFetchNewsPayload{},
				RetryTimes: 1,
			},
		}
	)

	for _, job := range jobs {
		entryID, err := w.queue.ScheduleTask(job.Task, job.Payload, job.CronSpec, job.RetryTimes)
		if err != nil {
			ctx.Logger().Error("error when initializing cronjob", err, appcontext.Fields{"job": job})
			panic(err)
		}

		ctx.Logger().Info(fmt.Sprintf("[cronjob] cronjob '%s' initialize successfully with cronSpec '%s' and retryTimes '%d'", job.Task, job.CronSpec, job.RetryTimes), appcontext.Fields{
			"entryId": entryID,
		})
	}
}
