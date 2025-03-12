package scraper

import "time"

type Article struct {
	Title       string
	Content     string
	URL         string
	Source      string
	Category    string
	PublishedAt time.Time
}
