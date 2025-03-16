package domain

import "github.com/namhq1989/go-utilities/appcontext"

type NlpRepository interface {
	ExtractWord(ctx *appcontext.AppContext, content string) (*NlpExtractWordResult, error)
	AnalyzeWord(ctx *appcontext.AppContext, word, category string) (*NlpAnalyzeWordResult, error)
}

type NlpExtractWordResult struct {
	Word     string
	Category string
}

type NlpWordExample struct {
	Level   Level
	Example string
	Word    string
}

type NlpWordInfo struct {
	Definitions []WordDefinition
	Ipa         string
	Level       Level
	Pos         []string
	NounForm    *WordNounForm
	VerbForm    *WordVerbForm
}

type NlpAnalyzeWordResult struct {
	Category string
	Examples []NlpWordExample
	Word     NlpWordInfo
}
