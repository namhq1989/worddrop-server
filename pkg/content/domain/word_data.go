package domain

type WordDefinition struct {
	Pos        string
	Definition string
}

type WordNounForm struct {
	Base   string
	Plural string
}

type WordVerbForm struct {
	Base               string
	Past               string
	PastParticiple     string
	Gerund             string
	PresentThirdPerson string
}
