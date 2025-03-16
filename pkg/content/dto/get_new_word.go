package dto

type GetNewWordRequest struct {
	Categories []string `json:"categories"`
	Level      string   `json:"level"`
}

type GetNewWordResponse struct {
	Word Word `json:"word"`
}
