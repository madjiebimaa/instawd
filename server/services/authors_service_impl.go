package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
	"github.com/madjiebimaa/go-random-quotes/models/web"
	"github.com/madjiebimaa/go-random-quotes/repositories"
)

type AuthorsServiceImpl struct {
	AuthorsRepository repositories.AuthorsRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAuthorsService(authorsRepository repositories.AuthorsRepository, db *sql.DB, validate *validator.Validate) AuthorsService {
	return &AuthorsServiceImpl{
		AuthorsRepository: authorsRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *AuthorsServiceImpl) Create(ctx context.Context, request web.AuthorCreateRequest) web.AuthorResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	id := helpers.RandomString(12)
	slug := helpers.ToSlugFromAuthorName(request.Name)

	author := domain.Author{
		Id:          id,
		Name:        request.Name,
		Link:        sql.NullString{String: request.Link, Valid: request.Link != ""},
		Bio:         sql.NullString{String: request.Bio, Valid: request.Bio != ""},
		Description: sql.NullString{String: request.Description, Valid: request.Description != ""},
		QuoteCount:  0,
		Slug:        slug,
	}

	author = service.AuthorsRepository.Create(ctx, tx, author)

	return helpers.ToAuthorResponse(author)
}

func (service *AuthorsServiceImpl) FindById(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	author := service.AuthorsRepository.FindById(ctx, tx, request.Id)

	return helpers.ToAuthorResponse(author)
}

func (service *AuthorsServiceImpl) FindAll(ctx context.Context) []web.AuthorResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	authors := service.AuthorsRepository.FindAll(ctx, tx)

	return helpers.ToAuthorResponses(authors)
}

func (service *AuthorsServiceImpl) FindBySlug(ctx context.Context, request web.AuthorFindBySlugRequest) web.AuthorResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	author := service.AuthorsRepository.FindBySlug(ctx, tx, request.Slug)

	return helpers.ToAuthorResponse(author)
}

func (service *AuthorsServiceImpl) FindAuthorAndQuotes(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorAndQuotesResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	author := service.AuthorsRepository.FindById(ctx, tx, request.Id)

	quoteRepository := repositories.NewQuotesRepository()
	quotes := quoteRepository.FindByAuthorId(ctx, tx, author.Id)

	return helpers.ToAuthorAndQuotesResponse(author, quotes)
}
