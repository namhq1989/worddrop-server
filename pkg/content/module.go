package content

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/monolith"
	"github.com/namhq1989/worddrop-server/pkg/content/infrastructure"
	"github.com/namhq1989/worddrop-server/pkg/content/worker"
)

type Module struct{}

func (Module) Name() string {
	return "CONTENT"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		wordRepository        = infrastructure.NewWordRepository(mono.Database())
		wordExampleRepository = infrastructure.NewWordExampleRepository(mono.Database())
		wordNewsRepository    = infrastructure.NewWordNewsRepository(mono.Database())

		externalApiRepository = infrastructure.NewExternalAPIRepository(mono.ExternalAPI())
		ttsRepository         = infrastructure.NewTTSRepository(mono.TTS())
		nlpRepository         = infrastructure.NewNlpRepository(mono.NLP())
		queueRepository       = infrastructure.NewQueueRepository(mono.Queue())
	)

	// worker
	w := worker.New(
		mono.Queue(),
		wordRepository,
		wordNewsRepository,
		wordExampleRepository,
		externalApiRepository,
		nlpRepository,
		queueRepository,
		ttsRepository,
	)
	w.Start()

	return nil
}
