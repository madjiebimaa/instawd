package web

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
