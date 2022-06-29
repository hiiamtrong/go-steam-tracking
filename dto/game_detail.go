package dto

type GetAllGameDetailRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type SearchGameDetailRequest struct {
	Query     string `json:"query"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	SortBy    string `json:"sort_by"`
	SortOrder int    `json:"sort_order"`
}

type GetGameDetailByIdRequest struct {
	Id int `json:"id"`
}
