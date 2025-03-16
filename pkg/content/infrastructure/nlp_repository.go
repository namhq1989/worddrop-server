package infrastructure

import (
	"errors"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/nlp"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type NlpRepository struct {
	nlp nlp.Operations
}

func NewNlpRepository(nlp nlp.Operations) NlpRepository {
	return NlpRepository{
		nlp: nlp,
	}
}

func (r NlpRepository) ExtractWord(ctx *appcontext.AppContext, content string) (*domain.NlpExtractWordResult, error) {
	apiResult, err := r.nlp.ExtractWord(ctx, content)
	if err != nil {
		return nil, err
	}
	if apiResult.Word == "" {
		return nil, errors.New("invalid_word")
	}

	return &domain.NlpExtractWordResult{
		Word:     apiResult.Word,
		Category: apiResult.Category,
	}, nil
}

func (r NlpRepository) AnalyzeWord(ctx *appcontext.AppContext, word, category string) (*domain.NlpAnalyzeWordResult, error) {
	apiResult, err := r.nlp.AnalyzeWord(ctx, word, category)
	if err != nil {
		return nil, err
	}

	examples := []domain.NlpWordExample{
		{
			Level:   domain.LevelBeginner,
			Example: apiResult.Examples.Beginner.Example,
			Word:    apiResult.Examples.Beginner.Word,
		},
		{
			Level:   domain.LevelIntermediate,
			Example: apiResult.Examples.Intermediate.Example,
			Word:    apiResult.Examples.Intermediate.Word,
		},
		{
			Level:   domain.LevelAdvanced,
			Example: apiResult.Examples.Advanced.Example,
			Word:    apiResult.Examples.Advanced.Word,
		},
	}

	definitions := make([]domain.WordDefinition, len(apiResult.Word.Definitions))
	for i, def := range apiResult.Word.Definitions {
		definitions[i] = domain.WordDefinition{
			Definition: def.Definition,
			Pos:        def.Pos,
		}
	}

	var nounForm *domain.WordNounForm
	if apiResult.Word.Noun != nil {
		nounForm = &domain.WordNounForm{
			Base:   apiResult.Word.Noun.Base,
			Plural: apiResult.Word.Noun.Plural,
		}
	}

	var verbForm *domain.WordVerbForm
	if apiResult.Word.Verb != nil {
		verbForm = &domain.WordVerbForm{
			Base:               apiResult.Word.Verb.Base,
			Gerund:             apiResult.Word.Verb.Gerund,
			Past:               apiResult.Word.Verb.Past,
			PastParticiple:     apiResult.Word.Verb.PastParticiple,
			PresentThirdPerson: apiResult.Word.Verb.PresentThirdPerson,
		}
	}

	// Create the result
	return &domain.NlpAnalyzeWordResult{
		Category: apiResult.Category,
		Examples: examples,
		Word: domain.NlpWordInfo{
			Definitions: definitions,
			Ipa:         apiResult.Word.Ipa,
			Level:       domain.ToLevel(apiResult.Word.Level),
			Pos:         apiResult.Word.Pos,
			NounForm:    nounForm,
			VerbForm:    verbForm,
		},
	}, nil
}
