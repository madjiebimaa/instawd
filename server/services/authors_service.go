package services

import (
	"context"

	"github.com/madjiebimaa/go-random-quotes/models/web"
)

type AuthorsService interface {
	Create(ctx context.Context, request web.AuthorCreateRequest) web.AuthorResponse
	FindById(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorResponse
	FindAll(ctx context.Context) []web.AuthorResponse
	FindBySlug(ctx context.Context, request web.AuthorFindBySlugRequest) web.AuthorResponse
	FindAuthorAndQuotes(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorAndQuotesResponse
}
