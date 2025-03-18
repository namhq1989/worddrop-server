package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/model"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type WordMapper struct{}

func (WordMapper) FromModelToDomain(word model.Words) (*domain.Word, error) {
	var result = &domain.Word{
		ID:            word.ID,
		Word:          word.Word,
		Level:         domain.ToLevel(word.Level),
		Definitions:   make([]domain.WordDefinition, 0),
		PartsOfSpeech: word.PartsOfSpeech,
		Ipa:           word.Ipa,
		NounForm:      nil,
		VerbForm:      nil,
		CreatedAt:     word.CreatedAt,
		LastFetchedAt: word.LastFetchedAt,
	}

	if word.Definitions != "" {
		if err := json.Unmarshal([]byte(word.Definitions), &result.Definitions); err != nil {
			return nil, err
		}
	}

	if word.NounForm != nil {
		if err := json.Unmarshal([]byte(*word.NounForm), &result.NounForm); err != nil {
			return nil, err
		}
	}
	if word.VerbForm != nil {
		if err := json.Unmarshal([]byte(*word.VerbForm), &result.VerbForm); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (WordMapper) FromDomainToModel(word domain.Word) (*model.Words, error) {
	var result = &model.Words{
		ID:            word.ID,
		Word:          word.Word,
		Level:         word.Level.String(),
		Definitions:   "",
		PartsOfSpeech: word.PartsOfSpeech,
		Ipa:           word.Ipa,
		CreatedAt:     word.CreatedAt,
		LastFetchedAt: word.LastFetchedAt,
	}

	definitions := make([]WordDefinition, 0)
	for _, definition := range word.Definitions {
		definitions = append(definitions, WordDefinition{
			Pos:        definition.Pos,
			Definition: definition.Definition,
		})
	}
	if data, err := json.Marshal(definitions); err != nil {
		return nil, err
	} else {
		result.Definitions = string(data)
	}

	if word.NounForm != nil {
		nounForm := WordNounForm{
			Base:   word.NounForm.Base,
			Plural: word.NounForm.Plural,
		}
		if data, err := json.Marshal(nounForm); err != nil {
			return nil, err
		} else {
			nounFormStr := string(data)
			result.NounForm = &nounFormStr
		}
	}

	if word.VerbForm != nil {
		verbForm := WordVerbForm{
			Base:               word.VerbForm.Base,
			Past:               word.VerbForm.Past,
			PastParticiple:     word.VerbForm.PastParticiple,
			Gerund:             word.VerbForm.Gerund,
			PresentThirdPerson: word.VerbForm.PresentThirdPerson,
		}
		if data, err := json.Marshal(verbForm); err != nil {
			return nil, err
		} else {
			verbFormStr := string(data)
			result.VerbForm = &verbFormStr
		}
	}

	return result, nil
}
