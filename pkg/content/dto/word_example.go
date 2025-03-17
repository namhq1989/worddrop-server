package dto

import "github.com/namhq1989/worddrop-server/pkg/content/domain"

type WordExample struct {
	ID       string `json:"id"`
	Example  string `json:"example"`
	MainWord string `json:"mainWord"`
	Level    string `json:"level"`
}

func (WordExample) FromDomain(example domain.WordExample) WordExample {
	return WordExample{
		ID:       example.ID,
		Example:  example.Example,
		MainWord: example.MainWord,
		Level:    example.Level.String(),
	}
}
