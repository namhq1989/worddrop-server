package mapping

import (
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/model"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type WordExampleMapper struct{}

func (WordExampleMapper) FromModelToDomain(example model.WordExamples) (*domain.WordExample, error) {
	var result = &domain.WordExample{
		ID:        example.ID,
		WordID:    example.WordID,
		Example:   example.Example,
		MainWord:  example.MainWord,
		Level:     domain.Level(example.Level),
		CreatedAt: example.CreatedAt,
	}

	return result, nil
}

func (WordExampleMapper) FromDomainToModel(example domain.WordExample) (*model.WordExamples, error) {
	var result = &model.WordExamples{
		ID:        example.ID,
		WordID:    example.WordID,
		Example:   example.Example,
		MainWord:  example.MainWord,
		Level:     example.Level.String(),
		CreatedAt: example.CreatedAt,
	}

	return result, nil
}
