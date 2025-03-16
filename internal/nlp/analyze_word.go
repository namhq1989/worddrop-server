package nlp

import "github.com/namhq1989/go-utilities/appcontext"

type WordExample struct {
	Example string `json:"example"`
	Word    string `json:"word"`
}

type WordExamples struct {
	Advanced     WordExample `json:"advanced"`
	Beginner     WordExample `json:"beginner"`
	Intermediate WordExample `json:"intermediate"`
}

type Definition struct {
	Definition string `json:"definition"`
	Pos        string `json:"pos"`
}

type NounForm struct {
	Base   string `json:"base"`
	Plural string `json:"plural"`
}

type VerbForm struct {
	Base               string `json:"base"`
	Gerund             string `json:"gerund"`
	Past               string `json:"past"`
	PastParticiple     string `json:"pastParticiple"`
	PresentThirdPerson string `json:"presentThirdPerson"`
}

type WordInfo struct {
	Definitions []Definition `json:"definitions"`
	Ipa         string       `json:"ipa"`
	Level       string       `json:"level"`
	Noun        *NounForm    `json:"noun,omitempty"`
	Pos         []string     `json:"pos"`
	Verb        *VerbForm    `json:"verb,omitempty"`
}

type AnalyzeWordResult struct {
	Category string       `json:"category"`
	Examples WordExamples `json:"examples"`
	Provider string       `json:"provider"`
	Word     WordInfo     `json:"word"`
}

func (n NLP) AnalyzeWord(_ *appcontext.AppContext, word string, category string) (result *AnalyzeWordResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"word":     word,
			"category": category,
		}).
		SetResult(&result).
		Post("/analyze-word")
	return
}
