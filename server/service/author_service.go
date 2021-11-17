package service

import (
	"context"

	"github.com/madjiebimaa/go-random-quotes/model/web"
)

type AuthorService interface {
	Create(ctx context.Context, request web.AuthorCreateRequest) web.AuthorResponse
	FindById(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorResponse
	FindAll(ctx context.Context) []web.AuthorResponse
}
