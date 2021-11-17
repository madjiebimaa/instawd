package web

type QuoteResponse struct {
	Id       string `json:"id"`
	AuthorId string `json:"author_id"`
	Content  string `json:"content"`
}
