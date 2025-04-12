package externalapi

import (
	"fmt"
	"net/url"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/worddrop-server/internal/utils/manipulation"
)

type FetchNewsWithNewsServiceResult struct {
	Articles []NewsArticle
}

type newsApiResponse struct {
	News []newsApiResponseCluster `json:"news"`
}

type newsApiResponseCluster struct {
	News []newsApiResponseArticle `json:"News"`
}

type newsApiResponseArticle struct {
	Title       string `json:"Title"`
	Url         string `json:"Url"`
	PublishedOn string `json:"PublishedOn"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	Summary     string `json:"Summary"`
}

func (ea ExternalAPI) FetchNewsWithNewsService(ctx *appcontext.AppContext) (*FetchNewsWithNewsServiceResult, error) {
	var (
		apiResults newsApiResponse
		apiKey     = ea.getRandomAPIKey()
	)

	_, err := ea.rapidApiNews.R().
		SetHeader("x-rapidapi-key", apiKey).
		SetQueryParams(map[string]string{
			"languages": "en",
		}).
		SetResult(&apiResults).
		Get("/v2/trending")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when fetching news with rapidapi", err, appcontext.Fields{})
		return nil, err
	}

	news, err := ea.parseNewsData(apiResults)
	if err != nil {
		ctx.Logger().Error("[externalapi] error when parsing news data", err, appcontext.Fields{})
		return nil, err
	}

	return &FetchNewsWithNewsServiceResult{
		Articles: news,
	}, nil
}

func (ea ExternalAPI) parseNewsData(data newsApiResponse) ([]NewsArticle, error) {
	var news = make([]NewsArticle, 0)
	for _, cluster := range data.News {
		for _, article := range cluster.News {
			if article.Description == "" {
				fmt.Println("Skipping article with empty description:", article.Title)
				continue
			}

			if containsNonLatinChars(article.Summary) {
				fmt.Println("Skipping article with non-Latin characters in summary:", article.Title)
				continue
			}

			publishedAt, err := time.Parse(time.RFC3339, article.PublishedOn)
			if err != nil {
				publishedAt = time.Now()
				fmt.Printf("Warning: Could not parse date %s: %v\n", article.PublishedOn, err)
			}

			a := NewsArticle{
				Title:       article.Title,
				Description: article.Description,
				Summary:     article.Summary,
				ImageURL:    article.Image,
				Url:         cleanURL(article.Url),
				PublishedAt: publishedAt,
			}

			news = append(news, a)
		}
	}

	return news, nil
}

func (ea ExternalAPI) getRandomAPIKey() string {
	if ea.rapidApiKeysLength == 1 {
		return ea.rapidApiKeys[0]
	}

	idx := manipulation.RandomIntInRange(0, ea.rapidApiKeysLength-1)
	return ea.rapidApiKeys[idx]
}

func cleanURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	parsedURL.RawQuery = ""
	return parsedURL.String()
}

func containsNonLatinChars(s string) bool {
	for _, r := range s {
		// Check if the character is outside the Latin character range
		// Basic Latin, Latin-1 Supplement, and Latin Extended ranges
		if r > 0x024F && // End of Latin Extended-B
			!(r >= 0x02B0 && r <= 0x02FF) && // Skip Spacing Modifier Letters
			!(r >= 0x2000 && r <= 0x206F) && // Skip General Punctuation
			!(r >= 0x0370 && r <= 0x03FF) { // Allow Greek for common characters
			return true
		}
	}
	return false
}
