package mapping

import (
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/model"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type WordNewsMapper struct{}

func (WordNewsMapper) FromModelToDomain(news model.WordNews) (*domain.WordNews, error) {
	var result = &domain.WordNews{
		ID:          news.ID,
		WordID:      news.WordID,
		Categories:  news.Categories,
		SourceURL:   news.SourceURL,
		Title:       news.Title,
		Summary:     news.Summary,
		ImageURL:    news.ImageURL,
		CreatedAt:   news.CreatedAt,
		PublishedAt: news.PublishedAt,
	}

	return result, nil
}

func (WordNewsMapper) FromDomainToModel(news domain.WordNews) (*model.WordNews, error) {
	var result = &model.WordNews{
		ID:          news.ID,
		WordID:      news.WordID,
		Categories:  news.Categories,
		SourceURL:   news.SourceURL,
		Title:       news.Title,
		Summary:     news.Summary,
		ImageURL:    news.ImageURL,
		CreatedAt:   news.CreatedAt,
		PublishedAt: news.PublishedAt,
	}

	return result, nil
}
