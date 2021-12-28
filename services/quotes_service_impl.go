package services

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/madjiebimaa/instawd/exceptions"
	"github.com/madjiebimaa/instawd/helpers"
	"github.com/madjiebimaa/instawd/models"
	"github.com/madjiebimaa/instawd/repositories"
)

type QuotesService struct {
	QuotesRepository *repositories.QuotesRepository
	DB               *sql.DB
}

func NewQuotesService(quotesRepository *repositories.QuotesRepository, db *sql.DB) *QuotesService {
	return &QuotesService{
		QuotesRepository: quotesRepository,
		DB:               db,
	}
}

func (s *QuotesService) Create(ctx context.Context, req models.QuoteCreateRequest) (*models.Quote, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quote := models.Quote{
		Id:       utils.UUIDv4(),
		AuthorId: req.AuthorId,
		Content:  req.Content,
	}

	err = s.QuotesRepository.Create(ctx, tx, &quote)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	return &quote, nil
}

func (s *QuotesService) FindById(ctx context.Context, req models.QuoteFindByIdRequest) (*models.Quote, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quote, err := s.QuotesRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if quote == nil {
		return nil, exceptions.ErrQuoteNotFound
	}

	return quote, nil
}

func (s *QuotesService) FindAll(ctx context.Context) ([]models.Quote, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quotes, err := s.QuotesRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if quotes == nil {
		return nil, exceptions.ErrQuoteNotFound
	}

	return quotes, err
}

func (s *QuotesService) FindRandom(ctx context.Context) (*models.Quote, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quote, err := s.QuotesRepository.FindRandom(ctx, tx)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	return quote, nil
}
