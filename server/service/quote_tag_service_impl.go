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

type QuoteTagServiceImpl struct {
	QuoteTagRepository repository.QuoteTagRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewQuoteTagService(quoteTagRepository repository.QuoteTagRepository, db *sql.DB, validate *validator.Validate) QuoteTagService {
	return &QuoteTagServiceImpl{
		QuoteTagRepository: quoteTagRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *QuoteTagServiceImpl) Create(ctx context.Context, request web.QuoteTagCreateRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	id := helper.RandomString(12)

	quoteTag := domain.QuoteTag{
		Id:   id,
		Name: request.Name,
	}

	quoteTag = service.QuoteTagRepository.Create(ctx, tx, quoteTag)

	return helper.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagServiceImpl) Delete(ctx context.Context, request web.QuoteTagDeleteRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quoteTag := service.QuoteTagRepository.FindById(ctx, tx, request.Id)
	// if quoteTag doesn't exist handle not created yet

	quoteTag = service.QuoteTagRepository.Delete(ctx, tx, quoteTag)

	return helper.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagServiceImpl) FindById(ctx context.Context, request web.QuoteTagFindByIdRequest) web.QuoteTagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quoteTag := service.QuoteTagRepository.FindById(ctx, tx, request.Id)

	return helper.ToQuoteTagResponse(quoteTag)
}

func (service *QuoteTagServiceImpl) FindAll(ctx context.Context) []web.QuoteTagResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quoteTags := service.QuoteTagRepository.FindAll(ctx, tx)

	return helper.ToQuoteTagResponses(quoteTags)
}
