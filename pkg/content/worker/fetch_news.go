package worker

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"go.opentelemetry.io/otel"
)

type FetchNewsHandler struct {
	wordRepository        domain.WordRepository
	wordNewsRepository    domain.WordNewsRepository
	wordExampleRepository domain.WordExampleRepository
	externalAPIRepository domain.ExternalAPIRepository
	nlpRepository         domain.NlpRepository
	queueRepository       domain.QueueRepository
}

func NewFetchNewsHandler(
	wordRepository domain.WordRepository,
	wordNewsRepository domain.WordNewsRepository,
	wordExampleRepository domain.WordExampleRepository,
	externalAPIRepository domain.ExternalAPIRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
) FetchNewsHandler {
	return FetchNewsHandler{
		wordRepository:        wordRepository,
		wordNewsRepository:    wordNewsRepository,
		wordExampleRepository: wordExampleRepository,
		externalAPIRepository: externalAPIRepository,
		nlpRepository:         nlpRepository,
		queueRepository:       queueRepository,
	}
}

func (h FetchNewsHandler) FetchNews(ctx *appcontext.AppContext, _ domain.QueueFetchNewsPayload) error {
	tracer := otel.Tracer("[tracer] fetch news")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] fetch news")
	ctx.SetContext(spanCtx)
	defer span.End()
	ctx.Logger().Text("fetch news from external api")
	news, err := h.externalAPIRepository.FetchNews(ctx)
	if err != nil {
		ctx.Logger().Error("failed to fetch news from external api", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("collect news urls")
	urls := make([]string, 0)
	for _, n := range news {
		urls = append(urls, n.Url)
	}

	ctx.Logger().Text("get news by urls")
	existedNews, err := h.wordNewsRepository.FindBySourceURLs(ctx, urls)
	if err != nil {
		ctx.Logger().Error("failed to get news by urls", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("filter the new news")
	newNews := make([]domain.NewsArticleScraped, 0)
	for _, n := range news {
		existed := false
		for _, en := range existedNews {
			if n.Url == en.SourceURL {
				existed = true
				break
			}
		}
		if !existed {
			newNews = append(newNews, n)
		}
	}

	ctx.Logger().Info("processing new news", appcontext.Fields{"count": len(newNews)})
	for _, article := range newNews {
		if err = h.processNewsArticle(ctx, article); err != nil {
			ctx.Logger().Error("failed to process news article", err, appcontext.Fields{"url": article.Url})
			continue
		}
	}

	return nil
}

func (h FetchNewsHandler) processNewsArticle(ctx *appcontext.AppContext, article domain.NewsArticleScraped) error {
	ctx.Logger().Info("call NLP server to extract word from article", appcontext.Fields{"url": article.Url})
	extractResult, err := h.nlpRepository.ExtractWord(ctx, article.Summary)
	if err != nil {
		ctx.Logger().Error("failed to extract word from article", err, appcontext.Fields{"url": article.Url})
		return err
	}

	ctx.Logger().Info("extracted word", appcontext.Fields{"word": extractResult.Word, "category": extractResult.Category})

	ctx.Logger().Text("find word in database")
	word, err := h.wordRepository.FindByWord(ctx, extractResult.Word)
	if err != nil {
		ctx.Logger().Error("failed to find word", err, appcontext.Fields{"word": extractResult.Word})
		return err
	}

	if word != nil {
		threeMonthsAgo := time.Now().AddDate(0, -3, 0)
		if word.LastFetchedAt.After(threeMonthsAgo) {
			ctx.Logger().Info("word was fetched recently, skipping",
				appcontext.Fields{
					"word":          extractResult.Word,
					"lastFetchedAt": word.LastFetchedAt,
					"reason":        "less than 3 months since last fetch",
				})
			return nil
		}
	}

	ctx.Logger().Info("call NLP server to analyzing word", appcontext.Fields{"word": extractResult.Word, "category": extractResult.Category})
	analysisResult, err := h.nlpRepository.AnalyzeWord(ctx, extractResult.Word, extractResult.Category)
	if err != nil {
		ctx.Logger().Error("failed to analyze word", err, appcontext.Fields{"word": extractResult.Word})
		return err
	}

	// Print the analysis response
	ctx.Logger().Info("word analysis complete", appcontext.Fields{
		"word":             extractResult.Word,
		"category":         analysisResult.Category,
		"level":            analysisResult.Word.Level,
		"pos":              analysisResult.Word.Pos,
		"definitionsCount": len(analysisResult.Word.Definitions),
		"examplesCount":    len(analysisResult.Examples),
	})

	if word == nil {
		ctx.Logger().Info("creating new word", appcontext.Fields{"word": extractResult.Word})
		word, err = domain.NewWord(
			extractResult.Word,
			string(analysisResult.Word.Level),
			analysisResult.Word.Definitions,
			analysisResult.Word.Pos,
			analysisResult.Word.Ipa,
			analysisResult.Word.NounForm,
			analysisResult.Word.VerbForm,
		)
		if err != nil {
			ctx.Logger().Error("failed to create new word", err, appcontext.Fields{"word": extractResult.Word})
			return err
		}

		ctx.Logger().Text("persist word to database")
		if err = h.wordRepository.Create(ctx, *word); err != nil {
			ctx.Logger().Error("failed to save word to database", err, appcontext.Fields{"word": word.Word})
			return err
		}
	} else {
		ctx.Logger().Info("updating existing word's last fetched time", appcontext.Fields{"word": word.Word, "id": word.ID})
		word.SetLastFetchedAt(manipulation.NowUTC())
		err = h.wordRepository.Update(ctx, *word)
		if err != nil {
			ctx.Logger().Error("failed to update word's last fetched time", err, appcontext.Fields{"word": word.Word, "id": word.ID})
			return err
		}
	}

	ctx.Logger().Info("adding job to generate word audio", appcontext.Fields{"word": extractResult.Word, "id": word.ID})
	err = h.queueRepository.GenerateTextAudio(ctx, domain.QueueGenerateTextAudioPayload{
		ID:      word.ID,
		Content: word.Word,
	})
	if err != nil {
		ctx.Logger().Error("failed to add job to generate word audio", err, appcontext.Fields{"word": extractResult.Word, "id": word.ID})
		return err
	}

	h.createWordExamples(ctx, *word, analysisResult.Examples)

	ctx.Logger().Info("creating news article", appcontext.Fields{"wordID": word.ID, "url": article.Url})
	news, err := domain.NewWordNews(
		word.ID,
		[]string{analysisResult.Category},
		article.Url,
		article.Title,
		article.Summary,
		article.ImageURL,
		article.PublishedAt,
	)
	if err != nil {
		ctx.Logger().Error("failed to create news article", err, appcontext.Fields{
			"wordID": word.ID,
			"url":    article.Url,
		})
		return err
	}

	ctx.Logger().Text("persist news article to database")
	err = h.wordNewsRepository.Create(ctx, *news)
	if err != nil {
		ctx.Logger().Error("failed to save news article to database", err, appcontext.Fields{
			"wordID": word.ID,
			"url":    article.Url,
		})
		return err
	}

	ctx.Logger().Info("word processing completed successfully", appcontext.Fields{
		"word":   extractResult.Word,
		"wordID": word.ID,
	})

	return nil
}

func (h FetchNewsHandler) createWordExamples(ctx *appcontext.AppContext, word domain.Word, examples []domain.NlpWordExample) {
	ctx.Logger().Info("creating word examples", appcontext.Fields{"count": len(examples)})
	for _, e := range examples {
		example, err := domain.NewWordExample(word.ID, e.Example, e.Word, e.Level.String())
		if err != nil {
			ctx.Logger().Error("failed to create word example", err, appcontext.Fields{
				"wordID":  word.ID,
				"example": e.Example,
			})
			continue
		}

		ctx.Logger().Text("persist word example to database")
		err = h.wordExampleRepository.Create(ctx, *example)
		if err != nil {
			ctx.Logger().Error("failed to save word example", err, appcontext.Fields{
				"wordID":  word.ID,
				"example": example.Example,
			})
			continue
		}

		ctx.Logger().Info("adding job to generate example audio", appcontext.Fields{
			"exampleID": example.ID,
			"example":   example.Example,
		})
		err = h.queueRepository.GenerateTextAudio(ctx, domain.QueueGenerateTextAudioPayload{
			ID:      example.ID,
			Content: example.Example,
		})

		if err != nil {
			ctx.Logger().Error("failed to add job to generate example audio", err, appcontext.Fields{
				"exampleID": example.ID,
				"example":   example.Example,
			})
			continue
		}
	}
}
