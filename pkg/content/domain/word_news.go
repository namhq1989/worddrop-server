package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

type WordNewsRepository interface {
	FindByWordID(ctx *appcontext.AppContext, wordID string) ([]WordNews, error)
	FindBySourceURLs(ctx *appcontext.AppContext, urls []string) ([]WordNews, error)
	Create(ctx *appcontext.AppContext, news WordNews) error
}

type WordNews struct {
	ID          string
	WordID      string
	Categories  []string
	SourceURL   string
	Title       string
	Summary     string
	ImageURL    string
	CreatedAt   time.Time
	PublishedAt time.Time
}

func NewWordNews(wordID string, categories []string, sourceURL string, title string, summary string, imageURL string, publishedAt time.Time) (*WordNews, error) {
	wn := WordNews{
		ID:          uuid.New(),
		Categories:  categories,
		ImageURL:    imageURL,
		CreatedAt:   manipulation.NowUTC(),
		PublishedAt: publishedAt,
	}

	if err := wn.SetWordID(wordID); err != nil {
		return nil, err
	}
	if err := wn.SetSourceURL(sourceURL); err != nil {
		return nil, err
	}
	if err := wn.SetTitle(title); err != nil {
		return nil, err
	}
	if err := wn.SetSummary(summary); err != nil {
		return nil, err
	}

	return &wn, nil
}

func (wn *WordNews) SetWordID(wordID string) error {
	if !uuid.IsValidID(wordID) {
		return apperrors.Common.InvalidID
	}

	wn.WordID = wordID
	return nil
}

func (wn *WordNews) SetTitle(title string) error {
	if title == "" {
		return apperrors.Common.InvalidNewsTitle
	}

	wn.Title = title
	return nil
}

func (wn *WordNews) SetSummary(summary string) error {
	if summary == "" {
		return apperrors.Common.InvalidNewsSummary
	}

	wn.Summary = summary
	return nil
}

func (wn *WordNews) SetSourceURL(sourceURL string) error {
	if sourceURL == "" {
		return apperrors.Common.InvalidNewsSourceURL
	}

	wn.SourceURL = sourceURL
	return nil
}
