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

type QuoteTagsService struct {
	QuoteTagsRepository *repositories.QuoteTagsRepository
	DB                  *sql.DB
}

func NewQuoteTagsService(quoteTagsRepository *repositories.QuoteTagsRepository, db *sql.DB) *QuoteTagsService {
	return &QuoteTagsService{
		QuoteTagsRepository: quoteTagsRepository,
		DB:                  db,
	}
}

func (s *QuoteTagsService) Create(ctx context.Context, req models.QuoteTagCreateRequest) (*models.QuoteTag, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quoteTag := models.QuoteTag{
		Id:   utils.UUIDv4(),
		Name: req.Name,
	}

	isQuoteTag, err := s.QuoteTagsRepository.FindByName(ctx, tx, quoteTag.Name)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if isQuoteTag != nil {
		return nil, exceptions.ErrQuoteTagAlreadyExist
	}

	err = s.QuoteTagsRepository.Create(ctx, tx, quoteTag)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	return &quoteTag, err
}

func (s *QuoteTagsService) FindAll(ctx context.Context) ([]models.QuoteTag, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	quoteTags, err := s.QuoteTagsRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if quoteTags == nil {
		return nil, exceptions.ErrQuoteTagNotFound
	}

	return quoteTags, err
}
