package models

// author entity

type Author struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
	QuoteCount  int    `json:"quote_count"`
	Slug        string `json:"slug"`
}

type AuthorAndQuotes struct {
	Author Author  `json:"author"`
	Quotes []Quote `json:"quotes"`
}

type AuthorAndQuote struct {
	Author Author `json:"author"`
	Quote  Quote  `json:"quote"`
}

// author requests

type AuthorCreateRequest struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
}

type AuthorFindByIdRequest struct {
	Id string `json:"id"`
}

type AuthorFindBySlugRequest struct {
	Slug string `json:"slug"`
}
