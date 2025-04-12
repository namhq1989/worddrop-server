package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

const (
	NewsService       = "news"
	GoogleNewsService = "google_news"
)

type ExternalAPIRepository interface {
	FetchNewsWithNewsService(ctx *appcontext.AppContext) ([]NewsArticleScraped, error)
	FetchNewsWithGoogleService(ctx *appcontext.AppContext) ([]NewsArticleScraped, error)
}

type NewsArticleScraped struct {
	Title       string
	Url         string
	Description string
	ImageURL    string
	Summary     string
	PublishedAt time.Time
}
