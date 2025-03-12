package nlp

import "github.com/namhq1989/go-utilities/appcontext"

type SummarizeNewsResult struct {
	Summary string                  `json:"summary"`
	Word    SummarizeNewsResultWord `json:"word"`
}

type SummarizeNewsResultWord struct {
	Word string                       `json:"word"`
	Ipa  string                       `json:"ipa"`
	Pos  string                       `json:"pos"`
	Noun *SummarizeNewsResultWordNoun `json:"noun"`
	Verb *SummarizeNewsResultWordVerb `json:"verb"`
}

type SummarizeNewsResultWordNoun struct {
	Base   string `json:"base"`
	Plural string `json:"plural"`
}

type SummarizeNewsResultWordVerb struct {
	Base               string `json:"base"`
	Past               string `json:"past"`
	PastParticiple     string `json:"pastParticiple"`
	Gerund             string `json:"gerund"`
	PresentThirdPerson string `json:"presentThirdPerson"`
}

func (n NLP) SummarizeNews(_ *appcontext.AppContext, content string) (result *SummarizeNewsResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"content": content,
		}).
		SetResult(&result).
		Post("/summarize-news")
	return
}
