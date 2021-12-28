package controllers

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/instawd/exceptions"
	"github.com/madjiebimaa/instawd/helpers"
	"github.com/madjiebimaa/instawd/models"
	"github.com/madjiebimaa/instawd/services"
)

type AuthorsControllerImpl struct {
	AuthorsService *services.AuthorsService
	Validate       *validator.Validate
}

func NewAuthorsController(authorsService *services.AuthorsService, validate *validator.Validate) *AuthorsControllerImpl {
	return &AuthorsControllerImpl{
		AuthorsService: authorsService,
		Validate:       validate,
	}
}

func (cr *AuthorsControllerImpl) Create(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	var req models.AuthorCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusBadRequest, "user input", exceptions.ErrInvalidInput)
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	author, err := cr.AuthorsService.Create(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusCreated, author)
}

func (cr *AuthorsControllerImpl) FindById(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	id := c.Params("authorId")
	req := models.AuthorFindByIdRequest{
		Id: id,
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	author, err := cr.AuthorsService.FindById(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, author)
}

func (cr *AuthorsControllerImpl) FindAll(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	authors, err := cr.AuthorsService.FindAll(ctx)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, authors)
}

func (cr *AuthorsControllerImpl) FindBySlug(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	slug := c.Params("authorSlug")
	req := models.AuthorFindBySlugRequest{
		Slug: slug,
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	author, err := cr.AuthorsService.FindBySlug(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, author)
}

// func (cr *AuthorsControllerImpl) FindAuthorAndQuotes(c *fiber.Ctx) error {
// 	now := time.Now()
// 	ctx := context.Background()

// 	id := c.Params("authorId")
// 	req := models.AuthorFindByIdRequest{
// 		Id: id,
// 	}

// 	author, err := cr.AuthorsService.Find(ctx, req)
// 	if err != nil {
// 		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
// 	}

// 	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, author)
// }
