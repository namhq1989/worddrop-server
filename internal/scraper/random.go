package scraper

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomWebsiteForCategory(category CategoryType) string {
	websites := websitesByCategory()[category]
	if len(websites) == 0 {
		return ""
	}

	randomIndex := rng.Intn(len(websites))
	return websites[randomIndex]
}

func randomWebsitesForAllCategories() map[CategoryType]string {
	result := make(map[CategoryType]string)

	for _, category := range AllCategories() {
		result[category] = randomWebsiteForCategory(category)
	}

	return result
}
