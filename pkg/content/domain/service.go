package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetWordExamples(ctx *appcontext.AppContext, wordID string) ([]WordExample, error)
	GetWordNews(ctx *appcontext.AppContext, wordID string) ([]WordNews, error)
}
