package scraper

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func newCollyCollector(domainName string, parallelism int, delay time.Duration) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domainName),
	)
	c.AllowURLRevisit = true
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  domainName,
		Parallelism: parallelism,
		Delay:       delay,
	})

	return c
}

func (s *Scraper) getCollectorForDomain(domain string) *colly.Collector {
	collector, exists := s.collectors[domain]
	if !exists {
		collector = newCollyCollector(domain, 10, 1*time.Second)
		s.collectors[domain] = collector
	}
	return collector
}
