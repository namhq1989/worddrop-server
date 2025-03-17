package domain

import "github.com/namhq1989/go-utilities/appcontext"

type CachingRepository interface {
	SetWordExamples(ctx *appcontext.AppContext, wordID string, examples []WordExample) error
	GetWordExamples(ctx *appcontext.AppContext, wordID string) ([]WordExample, error)

	SetWordNews(ctx *appcontext.AppContext, wordID string, news []WordNews) error
	GetWordNews(ctx *appcontext.AppContext, wordID string) ([]WordNews, error)
}
