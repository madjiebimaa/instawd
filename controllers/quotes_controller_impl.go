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

type QuoteController struct {
	QuotesService *services.QuotesService
	Validate      *validator.Validate
}

func NewQuoteController(quotesService *services.QuotesService, validate *validator.Validate) *QuoteController {
	return &QuoteController{
		QuotesService: quotesService,
		Validate:      validate,
	}
}

func (cr *QuoteController) Create(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	var req models.QuoteCreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusBadRequest, "user input", exceptions.ErrInvalidInput)
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	quote, err := cr.QuotesService.Create(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusCreated, quote)
}

func (cr *QuoteController) FindById(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	id := c.Params("quoteId")
	req := models.QuoteFindByIdRequest{
		Id: id,
	}

	if err := cr.Validate.Struct(req); err != nil {
		errors := helpers.ToErrorFields(err)
		return helpers.WebResponseFails(c, now, fiber.StatusBadRequest, errors)
	}

	quote, err := cr.QuotesService.FindById(ctx, req)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, quote)
}

// func (cr *QuoteController) FindQuoteDetail(c *fiber.Ctx) error {
// 	now := time.Now()
// 	ctx := context.Background()

// 	id := c.Params("quoteId")
// 	 req:= models.QuoteFindByIdRequest{
// 		 Id: id,
// 	 }

// 	quoteResponse := cr.QuotesService.FindQuoteAndAuthor(ctx, quote)

// 	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

// 	c.Status(fiber.StatusOK)
// 	c.Type(fiber.MIMEApplicationJSON)
// 	return c.JSON(webResponse)
// }

func (cr *QuoteController) FindAll(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	var filterRequest models.FilterRequest
	if err := c.QueryParser(&filterRequest); err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusBadRequest, "user input", exceptions.ErrInvalidInput)
	}

	ctx = context.WithValue(ctx, helpers.FILTER_REQUEST, filterRequest)
	quotes, err := cr.QuotesService.FindAll(ctx)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, quotes)
}

func (cr *QuoteController) FindRandom(c *fiber.Ctx) error {
	now := time.Now()
	ctx := context.Background()

	quote, err := cr.QuotesService.FindRandom(ctx)
	if err != nil {
		return helpers.WebResponseFail(c, now, fiber.StatusInternalServerError, "service", err)
	}

	return helpers.WebResponseSuccess(c, now, fiber.StatusOK, quote)
}

// func (cr *QuoteController) FindRandomAndAuthor(c *fiber.Ctx) error {
// 	now := time.Now()
// 	ctx := context.Background()

// 	ctx := context.Background()
// 	quoteResponse := controller.QuotesService.FindRandomAndAuthor(ctx)

// 	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

// 	c.Status(fiber.StatusOK)
// 	c.Type(fiber.MIMEApplicationJSON)
// 	return c.JSON(webResponse)
// }
