package nlp

import "github.com/namhq1989/go-utilities/appcontext"

type ExtractWordResult struct {
	Word     string `json:"word"`
	Category string `json:"category"`
}

func (n NLP) ExtractWord(_ *appcontext.AppContext, content string) (result *ExtractWordResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"content": content,
		}).
		SetResult(&result).
		Post("/extract-word")
	return
}
