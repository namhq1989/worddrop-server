package infrastructure

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/caching"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type CachingRepository struct {
	caching caching.Operations

	domain                  string
	wordExamplesCachingTime time.Duration
	wordNewsCachingTime     time.Duration
}

func NewCachingRepository(caching *caching.Caching, isEnvRelease bool) CachingRepository {
	if isEnvRelease {
		return CachingRepository{
			caching:                 caching,
			domain:                  "content",
			wordExamplesCachingTime: 1 * time.Hour,
			wordNewsCachingTime:     1 * time.Hour,
		}
	} else {
		cachingTime := 1 * time.Minute

		return CachingRepository{
			caching:                 caching,
			domain:                  "content",
			wordExamplesCachingTime: cachingTime,
			wordNewsCachingTime:     cachingTime,
		}
	}
}

// get word examples

func (r CachingRepository) GetWordExamples(ctx *appcontext.AppContext, wordID string) ([]domain.WordExample, error) {
	key := r.generateWordExamplesKey(wordID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result []domain.WordExample
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return result, nil
}

func (r CachingRepository) SetWordExamples(ctx *appcontext.AppContext, wordID string, examples []domain.WordExample) error {
	key := r.generateWordExamplesKey(wordID)
	r.caching.SetTTL(ctx, key, examples, r.wordExamplesCachingTime)
	return nil
}

func (r CachingRepository) generateWordExamplesKey(wordID string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("word:%s:examples", wordID))
}

// get word news

func (r CachingRepository) GetWordNews(ctx *appcontext.AppContext, wordID string) ([]domain.WordNews, error) {
	key := r.generateWordNewsKey(wordID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result []domain.WordNews
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return result, nil
}

func (r CachingRepository) SetWordNews(ctx *appcontext.AppContext, wordID string, news []domain.WordNews) error {
	key := r.generateWordNewsKey(wordID)
	r.caching.SetTTL(ctx, key, news, r.wordNewsCachingTime)
	return nil
}

func (r CachingRepository) generateWordNewsKey(wordID string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("word:%s:news", wordID))
}
