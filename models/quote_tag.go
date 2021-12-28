package models

// quote tag entity

type QuoteTag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// quote tag requests

type QuoteTagCreateRequest struct {
	Name string `json:"name"`
}

type QuoteTagFindByIdRequest struct {
	Id string `json:"id"`
}
