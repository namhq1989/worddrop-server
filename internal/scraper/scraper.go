package scraper

import (
	"time"

	"github.com/gocolly/colly/v2"
)

type Operations interface{}

type Scraper struct {
	collectors map[string]*colly.Collector
}

func NewScraper() *Scraper {
	s := &Scraper{
		collectors: make(map[string]*colly.Collector),
	}

	for _, websites := range websitesByCategory() {
		for _, domain := range websites {
			s.collectors[domain] = newCollyCollector(domain, 10, 1*time.Second)
		}
	}

	return s
}
