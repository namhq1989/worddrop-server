package dto

import (
	"github.com/namhq1989/worddrop-server/internal/utils/httprespond"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type WordNews struct {
	ID          string                    `json:"id"`
	Categories  []string                  `json:"categories"`
	SourceURL   string                    `json:"sourceUrl"`
	Title       string                    `json:"title"`
	Summary     string                    `json:"summary"`
	ImageURL    string                    `json:"imageUrl"`
	PublishedAt *httprespond.TimeResponse `json:"publishedAt"`
}

func (WordNews) FromDomain(news domain.WordNews) WordNews {
	return WordNews{
		ID:          news.ID,
		Categories:  news.Categories,
		SourceURL:   news.SourceURL,
		Title:       news.Title,
		Summary:     news.Summary,
		ImageURL:    news.ImageURL,
		PublishedAt: httprespond.NewTimeResponse(news.PublishedAt),
	}
}
