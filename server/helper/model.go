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

func ToQuoteNoAuthor(quote domain.Quote) web.QuoteNoAuthorResponse {
	return web.QuoteNoAuthorResponse{
		Id:      quote.Id,
		Content: quote.Content,
	}
}

func ToQuoteAndAuthorResponse(quote domain.Quote, author domain.Author) web.QuoteAndAuthorResponse {
	return web.QuoteAndAuthorResponse{
		Quote:  ToQuoteNoAuthor(quote),
		Author: ToAuthorResponse(author),
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
		Slug:        author.Slug,
	}
}

func ToAuthorResponses(authors []domain.Author) []web.AuthorResponse {
	var authorResponses []web.AuthorResponse
	for _, author := range authors {
		authorResponses = append(authorResponses, ToAuthorResponse(author))
	}

	return authorResponses
}

func ToNewWebResponse(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
