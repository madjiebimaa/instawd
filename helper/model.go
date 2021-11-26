package helper

import (
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

func ToNewWebResponse(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}

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

func ToQuoteNoAuthorResponse(quote domain.Quote) web.QuoteNoAuthorResponse {
	return web.QuoteNoAuthorResponse{
		Id:      quote.Id,
		Content: quote.Content,
	}
}

func ToQuoteNoAuthorResponses(quotes []domain.Quote) []web.QuoteNoAuthorResponse {
	var quoteNoAuthorResponses []web.QuoteNoAuthorResponse
	for _, quote := range quotes {
		quoteNoAuthorResponses = append(quoteNoAuthorResponses, ToQuoteNoAuthorResponse(quote))
	}

	return quoteNoAuthorResponses
}

func ToQuoteAndAuthorResponse(quote domain.Quote, author domain.Author) web.QuoteAndAuthorResponse {
	return web.QuoteAndAuthorResponse{
		Quote:  ToQuoteNoAuthorResponse(quote),
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

func ToAuthorAndQuotesResponse(author domain.Author, quotes []domain.Quote) web.AuthorAndQuotesResponse {
	return web.AuthorAndQuotesResponse{
		Author: ToAuthorResponse(author),
		Quotes: ToQuoteNoAuthorResponses(quotes),
	}
}

func ToQuoteTagResponse(quoteTag domain.QuoteTag) web.QuoteTagResponse {
	return web.QuoteTagResponse{
		Id:   quoteTag.Id,
		Name: quoteTag.Name,
	}
}

func ToQuoteTagResponses(quoteTags []domain.QuoteTag) []web.QuoteTagResponse {
	var quoteTagResponses []web.QuoteTagResponse
	for _, quoteTag := range quoteTags {
		quoteTagResponses = append(quoteTagResponses, ToQuoteTagResponse(quoteTag))
	}

	return quoteTagResponses
}

func ToQuoteRandomResponse(author domain.Author, quote domain.Quote) web.QuoteRandomResponse {
	return web.QuoteRandomResponse{
		Id:      quote.Id,
		Author:  author.Name,
		Content: quote.Content,
	}
}
