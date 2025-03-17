package content

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/monolith"
	"github.com/namhq1989/worddrop-server/pkg/content/application"
	"github.com/namhq1989/worddrop-server/pkg/content/infrastructure"
	"github.com/namhq1989/worddrop-server/pkg/content/rest"
	"github.com/namhq1989/worddrop-server/pkg/content/shared"
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
		cachingRepository     = infrastructure.NewCachingRepository(mono.Caching(), mono.Config().IsEnvRelease)

		service = shared.NewService(
			wordExampleRepository,
			wordNewsRepository,
			cachingRepository,
		)

		app = application.New(
			wordRepository,
			service,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

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
