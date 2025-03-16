package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

type WordExampleRepository interface {
	FindByWordID(ctx *appcontext.AppContext, wordID string) ([]WordExample, error)
	Create(ctx *appcontext.AppContext, example WordExample) error
}

type WordExample struct {
	ID        string
	WordID    string
	Example   string
	MainWord  string
	Level     Level
	CreatedAt time.Time
}

func NewWordExample(wordID, example, mainWord, level string) (*WordExample, error) {
	we := WordExample{
		ID:        uuid.New(),
		CreatedAt: manipulation.NowUTC(),
	}

	if err := we.SetWordID(wordID); err != nil {
		return nil, err
	}
	if err := we.SetExample(example); err != nil {
		return nil, err
	}
	if err := we.SetMainWord(mainWord); err != nil {
		return nil, err
	}
	if err := we.SetLevel(level); err != nil {
		return nil, err
	}

	return &we, nil
}

func (we *WordExample) SetWordID(wordID string) error {
	if !uuid.IsValidID(wordID) {
		return apperrors.Common.InvalidID
	}

	we.WordID = wordID
	return nil
}

func (we *WordExample) SetExample(example string) error {
	if example == "" {
		return apperrors.Common.InvalidExample
	}

	we.Example = example
	return nil
}

func (we *WordExample) SetMainWord(mainWord string) error {
	if mainWord == "" {
		return apperrors.Common.InvalidWord
	}

	we.MainWord = mainWord
	return nil
}

func (we *WordExample) SetLevel(level string) error {
	dLevel := ToLevel(level)
	if !dLevel.IsValid() {
		return apperrors.Common.InvalidLevel
	}

	we.Level = dLevel
	return nil
}
