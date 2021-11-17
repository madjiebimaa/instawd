package service

import (
	"context"

	"github.com/madjiebimaa/go-random-quotes/model/web"
)

type QuoteService interface {
	Create(ctx context.Context, request web.QuoteCreateRequest) web.QuoteResponse
	FindById(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteResponse
	FindAll(ctx context.Context) []web.QuoteResponse
	FindRandom(ctx context.Context) web.QuoteResponse
}
