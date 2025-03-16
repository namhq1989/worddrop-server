package domain

import (
	"time"

	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/internal/utils/pagetoken"
)

type WordFilter struct {
	Category  string
	Timestamp time.Time
	Limit     int64
}

func NewWordFilter(category string, pageToken string) (*WordFilter, error) {
	if category == "" {
		return nil, apperrors.Common.InvalidCategory
	}

	var (
		limit int64 = 10
		pt          = pagetoken.Decode(pageToken)
	)

	return &WordFilter{
		Category:  category,
		Timestamp: pt.Timestamp,
		Limit:     limit,
	}, nil
}
