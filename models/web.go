package models

// web requests

type SortRequest struct {
	ByName string `query:"by-name"`
}

type FilterRequest struct {
	Limit     int `query:"limit"`
	Offset    int `query:"offset"`
	MinLength int `query:"min-length"`
	MaxLength int `query:"max-length"`
}

// web responses

type FieldErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type WebResponse struct {
	Took    uint          `json:"took"`
	Success bool          `json:"success"`
	Status  int           `json:"status"`
	Data    interface{}   `json:"data"`
	Errors  []FieldErrors `json:"errors"`
}
