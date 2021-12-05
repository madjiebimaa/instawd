package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/web"
	"github.com/madjiebimaa/go-random-quotes/services"
)

type QuoteControllerImpl struct {
	QuotesService services.QuotesService
}

func NewQuoteController(quotesService services.QuotesService) QuotesController {
	return &QuoteControllerImpl{
		QuotesService: quotesService,
	}
}

func (controller *QuoteControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quote web.QuoteCreateRequest
	err := c.BodyParser(&quote)
	helpers.PanicIfError(err)

	ctx := context.Background()
	quoteResponse := controller.QuotesService.Create(ctx, quote)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindById(c *fiber.Ctx) error {
	var quote web.QuoteFindByIdRequest
	id := c.Params("quoteId")
	quote.Id = id

	ctx := context.Background()
	quoteResponse := controller.QuotesService.FindById(ctx, quote)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindQuoteAndAuthor(c *fiber.Ctx) error {
	var quote web.QuoteFindByIdRequest
	id := c.Params("quoteId")
	quote.Id = id

	ctx := context.Background()
	quoteResponse := controller.QuotesService.FindQuoteAndAuthor(ctx, quote)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindAll(c *fiber.Ctx) error {
	var filterRequest web.FilterRequest
	err := c.QueryParser(&filterRequest)
	helpers.PanicIfError(err)

	ctx := context.Background()
	ctx = context.WithValue(ctx, helpers.FILTER_REQUEST, filterRequest)

	quoteResponses := controller.QuotesService.FindAll(ctx)
	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindRandom(c *fiber.Ctx) error {
	ctx := context.Background()
	quoteResponse := controller.QuotesService.FindRandom(ctx)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindRandomAndAuthor(c *fiber.Ctx) error {
	ctx := context.Background()
	quoteResponse := controller.QuotesService.FindRandomAndAuthor(ctx)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}
