package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"github.com/namhq1989/worddrop-server/pkg/content/dto"
)

type GetInitialWordsHandler struct {
	wordRepository domain.WordRepository
	service        domain.Service
}

func NewGetInitialWordsHandler(wordRepository domain.WordRepository, service domain.Service) GetInitialWordsHandler {
	return GetInitialWordsHandler{
		wordRepository: wordRepository,
		service:        service,
	}
}

func (h GetInitialWordsHandler) GetInitialWords(ctx *appcontext.AppContext, _ dto.GetInitialWordsRequest) (*dto.GetInitialWordsResponse, error) {
	ctx.Logger().Text("new get initial words request")

	ctx.Logger().Text("create filter")
	filter, err := domain.NewWordFilter("", domain.InitialWordsCount)
	if err != nil {
		ctx.Logger().Error("failed to create filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find initial words")
	words, err := h.wordRepository.FindWithFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find initial words", err, appcontext.Fields{})
		return nil, err
	}
	if len(words) == 0 {
		ctx.Logger().Text("no initial words found, respond")
		return &dto.GetInitialWordsResponse{Words: make([]dto.Word, 0)}, nil
	}

	ctx.Logger().Text("find each word data")
	resp := make([]dto.Word, 0)
	for _, w := range words {
		word, wErr := h.findWordData(ctx, w)
		if wErr != nil {
			ctx.Logger().Error("failed to find word data", wErr, appcontext.Fields{"wordID": w.ID})
			continue
		}
		resp = append(resp, *word)
	}

	ctx.Logger().Text("done get initial words request")
	return &dto.GetInitialWordsResponse{Words: resp}, nil
}

func (h GetInitialWordsHandler) findWordData(ctx *appcontext.AppContext, word domain.Word) (*dto.Word, error) {
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
	resp := dto.Word{}.FromDomain(word, examples, news)
	return &resp, nil
}
