package helper

import (
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

func ToQuoteResponse(quote domain.Quote) web.QuoteResponse {
	return web.QuoteResponse{
		Id:       quote.AuthorId,
		Content:  quote.Content,
		AuthorId: quote.AuthorId,
	}
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
