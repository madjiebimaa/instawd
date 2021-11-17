package helper

import (
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

func ToQuoteResponse(quote domain.Quote) web.QuoteResponse {
	return web.QuoteResponse{
		Id:       quote.Id,
		AuthorId: quote.AuthorId,
		Content:  quote.Content,
	}
}

func ToQuoteResponses(quotes []domain.Quote) []web.QuoteResponse {
	var quoteResponses []web.QuoteResponse
	for _, quote := range quotes {
		quoteResponses = append(quoteResponses, ToQuoteResponse(quote))
	}

	return quoteResponses
}

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	return web.AuthorResponse{
		Id:          author.Id,
		Name:        author.Name,
		Link:        author.Link.String,
		Bio:         author.Bio.String,
		Description: author.Description.String,
		QuoteCount:  author.QuoteCount,
	}
}

func ToNewWebResponse(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
