package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
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
	return nil, nil
}
