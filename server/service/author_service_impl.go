package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
	"github.com/madjiebimaa/go-random-quotes/repository"
)

type AuthorServiceImpl struct {
	AuthorRepository repository.AuthorRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewAuthorService(authorRepository repository.AuthorRepository, db *sql.DB, validate *validator.Validate) AuthorService {
	return &AuthorServiceImpl{
		AuthorRepository: authorRepository,
		DB:               db,
		Validate:         validate,
	}
}

func (service *AuthorServiceImpl) Create(ctx context.Context, request web.AuthorCreateRequest) web.AuthorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	id := helper.RandomString(12)
	slug := helper.ToSlugFromAuthorName(request.Name)

	author := domain.Author{
		Id:          id,
		Name:        request.Name,
		Link:        sql.NullString{String: request.Link, Valid: request.Link != ""},
		Bio:         sql.NullString{String: request.Bio, Valid: request.Bio != ""},
		Description: sql.NullString{String: request.Description, Valid: request.Description != ""},
		QuoteCount:  0,
		Slug:        slug,
	}

	author = service.AuthorRepository.Create(ctx, tx, author)

	return helper.ToAuthorResponse(author)
}

func (service *AuthorServiceImpl) FindById(ctx context.Context, request web.AuthorFindByIdRequest) web.AuthorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	author := service.AuthorRepository.FindById(ctx, tx, request.Id)

	return helper.ToAuthorResponse(author)
}

func (service *AuthorServiceImpl) FindAll(ctx context.Context) []web.AuthorResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	authors := service.AuthorRepository.FindAll(ctx, tx)

	return helper.ToAuthorResponses(authors)
}
