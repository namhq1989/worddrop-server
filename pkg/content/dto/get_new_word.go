package dto

type GetNewWordRequest struct {
	Categories string `query:"categories"`
	Level      string `query:"level"`
}

type GetNewWordResponse struct {
	Word Word `json:"word"`
}
