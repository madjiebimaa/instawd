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

type AuthorsService struct {
	AuthorsRepository *repositories.AuthorsRepository
	DB                *sql.DB
}

func NewAuthorsService(authorsRepository *repositories.AuthorsRepository, db *sql.DB) *AuthorsService {
	return &AuthorsService{
		AuthorsRepository: authorsRepository,
		DB:                db,
	}
}

func (s *AuthorsService) Create(ctx context.Context, req models.AuthorCreateRequest) (*models.Author, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	id := utils.UUIDv4()
	slug := helpers.ToSlugFromAuthorName(req.Name)
	author := models.Author{
		Id:          id,
		Name:        req.Name,
		Link:        req.Link,
		Bio:         req.Bio,
		Description: req.Description,
		QuoteCount:  0,
		Slug:        slug,
	}

	err = s.AuthorsRepository.Create(ctx, tx, &author)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	return &author, nil
}

func (s *AuthorsService) FindById(ctx context.Context, req models.AuthorFindByIdRequest) (*models.Author, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	author, err := s.AuthorsRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if author == nil {
		return nil, exceptions.ErrAuthorNotFound
	}

	return author, err
}

func (s *AuthorsService) FindAll(ctx context.Context) ([]models.Author, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	authors, err := s.AuthorsRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if authors == nil {
		return nil, exceptions.ErrAuthorNotFound
	}

	return authors, nil
}

func (s *AuthorsService) FindBySlug(ctx context.Context, req models.AuthorFindBySlugRequest) (*models.Author, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}
	defer helpers.CommitOrRollBack(tx)

	author, err := s.AuthorsRepository.FindBySlug(ctx, tx, req.Slug)
	if err != nil {
		return nil, exceptions.ErrInternalServerError
	}

	if author == nil {
		return nil, exceptions.ErrAuthorNotFound
	}

	return author, nil
}
