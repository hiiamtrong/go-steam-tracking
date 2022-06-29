package dto

type GetAllGameDetailRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type SearchGameDetailRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
