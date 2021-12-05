package services

import (
	"context"

	"github.com/madjiebimaa/go-random-quotes/models/web"
)

type QuotesService interface {
	Create(ctx context.Context, request web.QuoteCreateRequest) web.QuoteResponse
	FindById(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteResponse
	FindQuoteAndAuthor(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteAndAuthorResponse
	FindAll(ctx context.Context) []web.QuoteResponse
	FindRandom(ctx context.Context) web.QuoteRandomResponse
	FindRandomAndAuthor(ctx context.Context) web.QuoteAndAuthorResponse
}
