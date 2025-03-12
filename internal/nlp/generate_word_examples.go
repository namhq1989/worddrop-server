package nlp

import "github.com/namhq1989/go-utilities/appcontext"

type GenerateWordExamplesResult struct {
	Word     string                             `json:"word"`
	Examples GenerateWordExamplesResultExamples `json:"examples"`
}

type GenerateWordExamplesResultExamples struct {
	Beginner     GenerateWordExamplesResultLevel `json:"beginner"`
	Intermediate GenerateWordExamplesResultLevel `json:"intermediate"`
	Advanced     GenerateWordExamplesResultLevel `json:"advanced"`
}

type GenerateWordExamplesResultLevel struct {
	Example string `json:"example"`
	Word    string `json:"word"`
}

func (n NLP) GenerateWordExamples(_ *appcontext.AppContext, word string) (result *GenerateWordExamplesResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"word": word,
		}).
		SetResult(&result).
		Post("/generate-examples")
	return
}
