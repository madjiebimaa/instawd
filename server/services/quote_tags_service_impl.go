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

type QuoteTagsServiceImpl struct {
	QuoteTagsRepository repositories.QuoteTagsRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewQuoteTagsService(quoteTagsRepository repositories.QuoteTagsRepository, db *sql.DB, validate *validator.Validate) QuoteTagsService {
	return &QuoteTagsServiceImpl{
		QuoteTagsRepository: quoteTagsRepository,
		DB:                  db,
		Validate:            validate,
	}
}

func (service *QuoteTagsServiceImpl) Create(ctx context.Context, request web.QuoteTagCreateRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	id := helpers.RandomString(12)

	quoteTag := domain.QuoteTag{
		Id:   id,
		Name: request.Name,
	}

	quoteTag = service.QuoteTagsRepository.Create(ctx, tx, quoteTag)

	return helpers.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagsServiceImpl) Delete(ctx context.Context, request web.QuoteTagDeleteRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quoteTag := service.QuoteTagsRepository.FindById(ctx, tx, request.Id)
	// if quoteTag doesn't exist handle not created yet

	quoteTag = service.QuoteTagsRepository.Delete(ctx, tx, quoteTag)

	return helpers.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagsServiceImpl) FindById(ctx context.Context, request web.QuoteTagFindByIdRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quoteTag := service.QuoteTagsRepository.FindById(ctx, tx, request.Id)

	return helpers.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagsServiceImpl) FindAll(ctx context.Context) []web.QuoteTagResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quoteTags := service.QuoteTagsRepository.FindAll(ctx, tx)

	return helpers.ToQuoteTagResponses(quoteTags)
}
