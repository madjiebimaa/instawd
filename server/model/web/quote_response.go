package web

type QuoteResponse struct {
	Id       string `json:"id"`
	AuthorId string `json:"author_id"`
	Content  string `json:"content"`
}

type QuoteNoAuthorResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type QuoteAndAuthorResponse struct {
	Quote  QuoteNoAuthorResponse `json:"quote"`
	Author AuthorResponse        `json:"author"`
}
