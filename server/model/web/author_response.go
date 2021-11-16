package web

type AuthorResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
	QuoteCount  int    `json:"quote_count"`
}
