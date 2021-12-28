package models

// quote entity

type Quote struct {
	Id       string `json:"id"`
	AuthorId string `json:"author_id"`
	Content  string `json:"content"`
}

// quote requests

type QuoteCreateRequest struct {
	AuthorId string `json:"author_id"`
	Content  string `json:"content"`
}

type QuoteFindByIdRequest struct {
	Id string `json:"id"`
}
