package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

func (s Service) GetWordExamples(ctx *appcontext.AppContext, wordID string) ([]domain.WordExample, error) {
	ctx.Logger().Info("[service] get word examples", appcontext.Fields{"wordID": wordID})

	ctx.Logger().Text("find word examples in caching")
	examples, _ := s.cachingRepository.GetWordExamples(ctx, wordID)
	if len(examples) > 0 {
		ctx.Logger().Text("word examples found in caching, return")
		return examples, nil
	}

	ctx.Logger().Text("word examples not found in caching, find in db")
	examples, err := s.wordExampleRepository.FindByWordID(ctx, wordID)
	if err != nil {
		ctx.Logger().Error("failed to find word examples in db", err, appcontext.Fields{})
		return nil, err
	}
	if len(examples) == 0 {
		ctx.Logger().ErrorText("word has no examples")
		return make([]domain.WordExample, 0), nil
	}

	ctx.Logger().Text("set word examples in caching")
	if err = s.cachingRepository.SetWordExamples(ctx, wordID, examples); err != nil {
		ctx.Logger().Error("failed to set word examples in caching", err, appcontext.Fields{})
		return nil, err
	}

	return examples, nil
}
