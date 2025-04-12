package externalapi

import (
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type FetchNewsWithGoogleServiceResult struct {
	Articles []NewsArticle
}

type googleNewsApiResponse struct {
	News []googleNewsArticle `json:"news"`
}

type googleNewsArticle struct {
	Date   string `json:"date"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	URL    string `json:"url"`
	Image  string `json:"image"`
	Source string `json:"source"`
}

func (ea ExternalAPI) FetchNewsWithGoogleService(ctx *appcontext.AppContext) (*FetchNewsWithGoogleServiceResult, error) {
	var (
		apiResults googleNewsApiResponse
		apiKey     = ea.getRandomAPIKey()
	)

	_, err := ea.rapidApiGoogleNews.R().
		SetHeader("x-rapidapi-key", apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"max_results": 25,
			"text":        "",
			"region":      "wt-wt",
		}).
		SetResult(&apiResults).
		Post("/")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when fetching news with Google News API", err, appcontext.Fields{})
		return nil, err
	}

	news, err := ea.parseGoogleNewsData(apiResults)
	if err != nil {
		ctx.Logger().Error("[externalapi] error when parsing Google News data", err, appcontext.Fields{})
		return nil, err
	}

	return &FetchNewsWithGoogleServiceResult{
		Articles: news,
	}, nil
}

func (ea ExternalAPI) parseGoogleNewsData(data googleNewsApiResponse) ([]NewsArticle, error) {
	var news = make([]NewsArticle, 0)
	for _, article := range data.News {
		if article.Body == "" {
			fmt.Println("Skipping article with empty body:", article.Title)
			continue
		}

		if containsNonLatinChars(article.Body) {
			fmt.Println("Skipping article with non-Latin characters in body:", article.Title)
			continue
		}

		publishedAt, err := time.Parse(time.RFC3339, article.Date)
		if err != nil {
			publishedAt = time.Now()
			fmt.Printf("Warning: Could not parse date %s: %v\n", article.Date, err)
		}

		a := NewsArticle{
			Title:       article.Title,
			Description: article.Body,
			Summary:     article.Body,
			ImageURL:    article.Image,
			Url:         cleanURL(article.URL),
			PublishedAt: publishedAt,
		}

		news = append(news, a)
	}

	return news, nil
}
