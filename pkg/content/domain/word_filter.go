package domain

import (
	"time"

	"github.com/namhq1989/worddrop-server/internal/utils/pagetoken"
)

type WordFilter struct {
	Timestamp time.Time
	Limit     int64
}

func NewWordFilter(pageToken string, limit int64) (*WordFilter, error) {
	pt := pagetoken.Decode(pageToken)
	return &WordFilter{
		Timestamp: pt.Timestamp,
		Limit:     limit,
	}, nil
}
