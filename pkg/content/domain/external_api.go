package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type ExternalAPIRepository interface {
	FetchNews(ctx *appcontext.AppContext) ([]NewsArticleScraped, error)
}

type NewsArticleScraped struct {
	Title       string
	Url         string
	Description string
	ImageURL    string
	Summary     string
	PublishedAt time.Time
}
