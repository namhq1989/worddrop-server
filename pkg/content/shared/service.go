package shared

import "github.com/namhq1989/worddrop-server/pkg/content/domain"

type Service struct {
	wordExampleRepository domain.WordExampleRepository
	wordNewsRepository    domain.WordNewsRepository
	cachingRepository     domain.CachingRepository
}

func NewService(
	wordExampleRepository domain.WordExampleRepository,
	wordNewsRepository domain.WordNewsRepository,
	cachingRepository domain.CachingRepository,
) Service {
	return Service{
		wordExampleRepository: wordExampleRepository,
		wordNewsRepository:    wordNewsRepository,
		cachingRepository:     cachingRepository,
	}
}
