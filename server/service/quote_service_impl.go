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

type QuoteServiceImpl struct {
	QuoteRepository repository.QuoteRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewQuoteService(quoteRepository repository.QuoteRepository, db *sql.DB, validate *validator.Validate) QuoteService {
	return &QuoteServiceImpl{
		QuoteRepository: quoteRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *QuoteServiceImpl) Create(ctx context.Context, request web.QuoteCreateRequest) web.QuoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	id := helper.RandomString(12)

	quote := domain.Quote{
		Id:       id,
		AuthorId: request.AuthorId,
		Content:  request.Content,
	}

	quote = service.QuoteRepository.Create(ctx, tx, quote)

	return helper.ToQuoteResponse(quote)
}

func (service *QuoteServiceImpl) FindById(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quote := service.QuoteRepository.FindById(ctx, tx, request.Id)

	return helper.ToQuoteResponse(quote)
}

func (service *QuoteServiceImpl) FindQuoteAndAuthor(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteAndAuthorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quote := service.QuoteRepository.FindById(ctx, tx, request.Id)

	authorRepository := repository.NewAuthorRepository()
	author := authorRepository.FindById(ctx, tx, quote.AuthorId)

	return helper.ToQuoteAndAuthorResponse(quote, author)
}

func (service *QuoteServiceImpl) FindAll(ctx context.Context) []web.QuoteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quotes := service.QuoteRepository.FindAll(ctx, tx)

	return helper.ToQuoteResponses(quotes)
}

func (service *QuoteServiceImpl) FindRandom(ctx context.Context) web.QuoteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	quote := service.QuoteRepository.FindRandom(ctx, tx)

	return helper.ToQuoteResponse(quote)
}
