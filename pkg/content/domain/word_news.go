package domain

import "time"

type WordNews struct {
	ID         string
	WordID     string
	Categories []string
	SourceURL  string
	Title      string
	Summary    string
	CreatedAt  time.Time
}
