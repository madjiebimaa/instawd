package web

type QuoteCreateRequest struct {
	AuthorId string `json:"author_id"`
	Content  string `json:"content"`
}

type QuoteFindByIdRequest struct {
	Id string `json:"id"`
}
