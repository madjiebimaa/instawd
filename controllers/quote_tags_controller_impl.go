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

type QuoteTagsController struct {
	QuoteTagsService *services.QuoteTagsService
	Validate         *validator.Validate
}

func NewQuoteTagController(quoteTagsService *services.QuoteTagsService, validate *validator.Validate) *QuoteTagsController {
	return &QuoteTagsController{
		QuoteTagsService: quoteTagsService,
		Validate:         validate,
	}
}

func (cr *QuoteTagsController) Create(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	var req models.QuoteTagCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusBadRequest, "user input", exceptions.ErrInvalidInput)
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	quoteTag, err := cr.QuoteTagsService.Create(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusCreated, quoteTag)
}

func (cr *QuoteTagsController) FindAll(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	quoteTags, err := cr.QuoteTagsService.FindAll(ctx)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, quoteTags)
}
