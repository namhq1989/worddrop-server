package mapping

type WordDefinition struct {
	Pos        string `json:"pos"`
	Definition string `json:"definition"`
}

type WordNounForm struct {
	Base   string `json:"base"`
	Plural string `json:"plural"`
}

type WordVerbForm struct {
	Base               string `json:"base"`
	Past               string `json:"past"`
	PastParticiple     string `json:"pastParticiple"`
	Gerund             string `json:"gerund"`
	PresentThirdPerson string `json:"presentThirdPerson"`
}
