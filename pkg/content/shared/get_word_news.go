package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

func (s Service) GetWordNews(ctx *appcontext.AppContext, wordID string) ([]domain.WordNews, error) {
	ctx.Logger().Info("[service] get word news", appcontext.Fields{"wordID": wordID})

	ctx.Logger().Text("find word news in caching")
	news, _ := s.cachingRepository.GetWordNews(ctx, wordID)
	if len(news) > 0 {
		ctx.Logger().Text("word news found in caching, return")
		return news, nil
	}

	ctx.Logger().Text("word news not found in caching, find in db")
	news, err := s.wordNewsRepository.FindByWordID(ctx, wordID)
	if err != nil {
		ctx.Logger().Error("failed to find word news in db", err, appcontext.Fields{})
		return nil, err
	}
	if len(news) == 0 {
		ctx.Logger().ErrorText("word has no news")
		return make([]domain.WordNews, 0), nil
	}

	ctx.Logger().Text("set word news in caching")
	if err = s.cachingRepository.SetWordNews(ctx, wordID, news); err != nil {
		ctx.Logger().Error("failed to set word news in caching", err, appcontext.Fields{})
		return nil, err
	}

	return news, nil
}
