package dto

import (
	"github.com/namhq1989/worddrop-server/internal/utils/httprespond"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
)

type Word struct {
	ID            string                    `json:"id"`
	Word          string                    `json:"word"`
	Level         string                    `json:"level"`
	Definitions   []WordDefinition          `json:"definitions"`
	PartsOfSpeech []string                  `json:"partsOfSpeech"`
	Ipa           string                    `json:"ipa"`
	NounForm      *WordNounForm             `json:"nounForm"`
	VerbForm      *WordVerbForm             `json:"verbForm"`
	Examples      []WordExample             `json:"examples"`
	News          []WordNews                `json:"news"`
	LastFetchedAt *httprespond.TimeResponse `json:"date"`
}

func (Word) FromDomain(word domain.Word, examples []domain.WordExample, news []domain.WordNews) Word {
	w := Word{
		ID:            word.ID,
		Word:          word.Word,
		Level:         word.Level.String(),
		Definitions:   make([]WordDefinition, 0),
		PartsOfSpeech: word.PartsOfSpeech,
		Ipa:           word.Ipa,
		NounForm:      nil,
		VerbForm:      nil,
		Examples:      make([]WordExample, 0),
		News:          make([]WordNews, 0),
		LastFetchedAt: httprespond.NewTimeResponse(word.LastFetchedAt),
	}

	for _, definition := range word.Definitions {
		w.Definitions = append(w.Definitions, WordDefinition{
			Pos:        definition.Pos,
			Definition: definition.Definition,
		})
	}
	if word.NounForm != nil {
		w.NounForm = &WordNounForm{
			Base:   word.NounForm.Base,
			Plural: word.NounForm.Plural,
		}
	}
	if word.VerbForm != nil {
		w.VerbForm = &WordVerbForm{
			Base:               word.VerbForm.Base,
			Past:               word.VerbForm.Past,
			PastParticiple:     word.VerbForm.PastParticiple,
			Gerund:             word.VerbForm.Gerund,
			PresentThirdPerson: word.VerbForm.PresentThirdPerson,
		}
	}

	for _, e := range examples {
		w.Examples = append(w.Examples, WordExample{}.FromDomain(e))
	}
	for _, n := range news {
		w.News = append(w.News, WordNews{}.FromDomain(n))
	}

	return w
}
