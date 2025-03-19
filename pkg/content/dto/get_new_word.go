package dto

type GetNewWordRequest struct {
	Categories string `query:"categories"`
	Levels     string `query:"levels"`
}

type GetNewWordResponse struct {
	Word Word `json:"word"`
}
