package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

const (
	NewWordTimeThreshold = 24 * 2 * time.Hour
)

type WordRepository interface {
	FindWithFilter(ctx *appcontext.AppContext, filter WordFilter) ([]Word, error)
	FindNewWord(ctx *appcontext.AppContext, categories []string, levels []string, ts time.Time) (*Word, error)
	FindByWord(ctx *appcontext.AppContext, word string) (*Word, error)
	FindByID(ctx *appcontext.AppContext, wordID string) (*Word, error)
	FindSimilar(ctx *appcontext.AppContext, pos, level string, ts time.Time, limit int64) ([]Word, error)
	Create(ctx *appcontext.AppContext, word Word) error
	Update(ctx *appcontext.AppContext, word Word) error
	Delete(ctx *appcontext.AppContext, word Word) error
}

type Word struct {
	ID            string
	Word          string
	Level         Level
	Definitions   []WordDefinition
	PartsOfSpeech []string
	Ipa           string
	NounForm      *WordNounForm
	VerbForm      *WordVerbForm
	CreatedAt     time.Time
	LastFetchedAt time.Time
}

func NewWord(word, level string, definitions []WordDefinition, partsOfSpeech []string, ipa string, nounForm *WordNounForm, verbForm *WordVerbForm) (*Word, error) {
	w := Word{
		ID:            uuid.New(),
		Definitions:   make([]WordDefinition, 0),
		PartsOfSpeech: make([]string, 0),
		CreatedAt:     manipulation.NowUTC(),
		LastFetchedAt: manipulation.NowUTC(),
	}

	if err := w.SetWord(word); err != nil {
		return nil, err
	}
	if err := w.SetLevel(level); err != nil {
		return nil, err
	}
	if err := w.SetDefinitions(definitions); err != nil {
		return nil, err
	}
	if err := w.SetPartsOfSpeech(partsOfSpeech); err != nil {
		return nil, err
	}
	if err := w.SetIpa(ipa); err != nil {
		return nil, err
	}
	if err := w.SetNounForm(nounForm); err != nil {
		return nil, err
	}
	if err := w.SetVerbForm(verbForm); err != nil {
		return nil, err
	}

	return &w, nil
}

func (w *Word) SetWord(word string) error {
	if word == "" {
		return apperrors.Common.InvalidWord
	}

	w.Word = word
	return nil
}

func (w *Word) SetLevel(level string) error {
	dLevel := ToLevel(level)
	if !dLevel.IsValid() {
		return apperrors.Common.InvalidLevel
	}

	w.Level = dLevel
	return nil
}

func (w *Word) SetDefinitions(definitions []WordDefinition) error {
	if len(definitions) == 0 {
		return apperrors.Common.InvalidDefinition
	}

	w.Definitions = definitions
	return nil
}

func (w *Word) SetPartsOfSpeech(partsOfSpeech []string) error {
	if len(partsOfSpeech) == 0 {
		return apperrors.Common.InvalidPartOfSpeech
	}

	w.PartsOfSpeech = partsOfSpeech
	return nil
}

func (w *Word) SetIpa(ipa string) error {
	if ipa == "" {
		return apperrors.Common.InvalidIPA
	}

	w.Ipa = ipa
	return nil
}

func (w *Word) SetNounForm(nounForm *WordNounForm) error {
	w.NounForm = nounForm
	return nil
}

func (w *Word) SetVerbForm(verbForm *WordVerbForm) error {
	w.VerbForm = verbForm
	return nil
}

func (w *Word) SetLastFetchedAt(lastFetchedAt time.Time) {
	w.LastFetchedAt = lastFetchedAt
}
