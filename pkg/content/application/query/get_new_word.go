package query

import (
	"strings"

	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"github.com/namhq1989/worddrop-server/pkg/content/dto"
)

type GetNewWordHandler struct {
	wordRepository domain.WordRepository
	service        domain.Service
}

func NewGetNewWordHandler(wordRepository domain.WordRepository, service domain.Service) GetNewWordHandler {
	return GetNewWordHandler{
		wordRepository: wordRepository,
		service:        service,
	}
}

func (h GetNewWordHandler) GetNewWord(ctx *appcontext.AppContext, req dto.GetNewWordRequest) (*dto.GetNewWordResponse, error) {
	ctx.Logger().Info("new get new word request", appcontext.Fields{"categories": req.Categories, "levels": req.Levels})
	categories := strings.Split(req.Categories, ",")
	levels := strings.Split(req.Levels, ",")

	ctx.Logger().Text("find new word in database")
	ts := manipulation.NowUTC().Add(domain.NewWordTimeThreshold * -1)
	word, err := h.wordRepository.FindNewWord(ctx, categories, levels, ts)
	if err != nil {
		ctx.Logger().Error("failed to find new word in database", err, appcontext.Fields{})
		return nil, err
	}
	if word == nil {
		ctx.Logger().ErrorText("no new word found")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("find word examples")
	examples, err := h.service.GetWordExamples(ctx, word.ID)
	if err != nil {
		ctx.Logger().Error("failed to find word examples", err, appcontext.Fields{"wordID": word.ID})
		examples = make([]domain.WordExample, 0)
	}

	ctx.Logger().Text("find word news")
	news, err := h.service.GetWordNews(ctx, word.ID)
	if err != nil {
		ctx.Logger().Error("failed to find word news", err, appcontext.Fields{"wordID": word.ID})
		news = make([]domain.WordNews, 0)
	}

	ctx.Logger().Text("convert to response")
	resp := dto.Word{}.FromDomain(*word, examples, news)

	ctx.Logger().Text("done get new word request")
	return &dto.GetNewWordResponse{
		Word: resp,
	}, nil
}
