package web

type FilterRequest struct {
	Limit     int `query:"limit"`
	Offset    int `query:"offset"`
	MinLength int `query:"min-length"`
	MaxLength int `query:"max-length"`
}
