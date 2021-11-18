package web

type QuoteTagCreateRequest struct {
	Name string `json:"name"`
}

type QuoteTagDeleteRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type QuoteTagFindByIdRequest struct {
	Id string `json:"id"`
}
