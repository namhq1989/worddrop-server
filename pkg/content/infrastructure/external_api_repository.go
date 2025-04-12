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

func (r ExternalAPIRepository) mapArticlesToDomain(articles []externalapi.NewsArticle) []domain.NewsArticleScraped {
	var result = make([]domain.NewsArticleScraped, 0, len(articles))
	for _, article := range articles {
		result = append(result, domain.NewsArticleScraped{
			Title:       article.Title,
			Url:         article.Url,
			Description: article.Description,
			ImageURL:    article.ImageURL,
			Summary:     article.Summary,
			PublishedAt: article.PublishedAt,
		})
	}

	return result
}

func (r ExternalAPIRepository) FetchNewsWithNewsService(ctx *appcontext.AppContext) ([]domain.NewsArticleScraped, error) {
	apiResult, err := r.ea.FetchNewsWithNewsService(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapArticlesToDomain(apiResult.Articles), nil
}

func (r ExternalAPIRepository) FetchNewsWithGoogleService(ctx *appcontext.AppContext) ([]domain.NewsArticleScraped, error) {
	apiResult, err := r.ea.FetchNewsWithGoogleService(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapArticlesToDomain(apiResult.Articles), nil
}
