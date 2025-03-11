package domain

import "time"

type Word struct {
	ID            string
	Word          string
	Definitions   []WordDefinition
	PartsOfSpeech []PartOfSpeech
	Ipa           string
	Audio         string
	CreatedAt     time.Time
}
