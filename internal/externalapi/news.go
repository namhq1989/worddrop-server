package externalapi

import "time"

type NewsArticle struct {
	Title       string
	Url         string
	Description string
	ImageURL    string
	Summary     string
	PublishedAt time.Time
}
