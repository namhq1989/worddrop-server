package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/pkg/content/application/query"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"github.com/namhq1989/worddrop-server/pkg/content/dto"
)

type (
	Queries interface {
		GetNewWord(ctx *appcontext.AppContext, req dto.GetNewWordRequest) (*dto.GetNewWordResponse, error)
	}
	Instance interface {
		Queries
	}

	queryHandlers struct {
		query.GetNewWordHandler
	}
	Application struct {
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	wordRepository domain.WordRepository,
	service domain.Service,
) *Application {
	return &Application{
		queryHandlers: queryHandlers{
			GetNewWordHandler: query.NewGetNewWordHandler(wordRepository, service),
		},
	}
}
