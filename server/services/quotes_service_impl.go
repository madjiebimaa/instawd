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

type QuotesServiceImpl struct {
	QuotesRepository repositories.QuotesRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewQuotesService(quotesRepository repositories.QuotesRepository, db *sql.DB, validate *validator.Validate) QuotesService {
	return &QuotesServiceImpl{
		QuotesRepository: quotesRepository,
		DB:               db,
		Validate:         validate,
	}
}

func (service *QuotesServiceImpl) Create(ctx context.Context, request web.QuoteCreateRequest) web.QuoteResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	id := helpers.RandomString(12)

	quote := domain.Quote{
		Id:       id,
		AuthorId: request.AuthorId,
		Content:  request.Content,
	}

	quote = service.QuotesRepository.Create(ctx, tx, quote)

	return helpers.ToQuoteResponse(quote)
}

func (service *QuotesServiceImpl) FindById(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quote := service.QuotesRepository.FindById(ctx, tx, request.Id)

	return helpers.ToQuoteResponse(quote)
}

func (service *QuotesServiceImpl) FindQuoteAndAuthor(ctx context.Context, request web.QuoteFindByIdRequest) web.QuoteAndAuthorResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quote := service.QuotesRepository.FindById(ctx, tx, request.Id)

	authorRepository := repositories.NewAuthorsRepository()
	author := authorRepository.FindById(ctx, tx, quote.AuthorId)

	return helpers.ToQuoteAndAuthorResponse(quote, author)
}

func (service *QuotesServiceImpl) FindAll(ctx context.Context) []web.QuoteResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quotes := service.QuotesRepository.FindAll(ctx, tx)

	return helpers.ToQuoteResponses(quotes)
}

func (service *QuotesServiceImpl) FindRandom(ctx context.Context) web.QuoteRandomResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quote := service.QuotesRepository.FindRandom(ctx, tx)

	authorRepository := repositories.NewAuthorsRepository()
	author := authorRepository.FindById(ctx, tx, quote.AuthorId)

	return helpers.ToQuoteRandomResponse(author, quote)
}

func (service *QuotesServiceImpl) FindRandomAndAuthor(ctx context.Context) web.QuoteAndAuthorResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollBack(tx)

	quote := service.QuotesRepository.FindRandom(ctx, tx)

	authorRepository := repositories.NewAuthorsRepository()
	author := authorRepository.FindById(ctx, tx, quote.AuthorId)

	return helpers.ToQuoteAndAuthorResponse(quote, author)
}
