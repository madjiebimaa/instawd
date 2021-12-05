package services

import (
	"context"

	"github.com/madjiebimaa/go-random-quotes/models/web"
)

type QuoteTagsService interface {
	Create(ctx context.Context, request web.QuoteTagCreateRequest) web.QuoteTagResponse
	Delete(ctx context.Context, request web.QuoteTagDeleteRequest) web.QuoteTagResponse
	FindById(ctx context.Context, request web.QuoteTagFindByIdRequest) web.QuoteTagResponse
	FindAll(ctx context.Context) []web.QuoteTagResponse
}
