package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/externalapi"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type ExternalAPIRepository struct {
	ea externalapi.Operations
}

func NewExternalAPIRepository(ea externalapi.Operations) ExternalAPIRepository {
	return ExternalAPIRepository{
		ea: ea,
	}
}

func (r ExternalAPIRepository) FetchNews(ctx *appcontext.AppContext) ([]domain.NewsArticleScraped, error) {
	apiResult, err := r.ea.FetchNews(ctx)
	if err != nil {
		return nil, err
	}

	var result = make([]domain.NewsArticleScraped, 0)
	for _, article := range apiResult.Articles {
		result = append(result, domain.NewsArticleScraped{
			Title:       article.Title,
			Url:         article.Url,
			Description: article.Description,
			ImageURL:    article.ImageURL,
			Summary:     article.Summary,
			PublishedAt: article.PublishedAt,
		})
	}

	return result, nil
}
