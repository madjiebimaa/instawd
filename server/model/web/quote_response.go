package web

type QuoteResponse struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	AuthorId string `json:"author_id"`
}
