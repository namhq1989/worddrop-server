package dto

type GetInitialWordsRequest struct{}

type GetInitialWordsResponse struct {
	Words []Word `json:"words"`
}
