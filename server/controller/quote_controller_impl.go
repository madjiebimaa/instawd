package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/web"
	"github.com/madjiebimaa/go-random-quotes/service"
)

type QuoteControllerImpl struct {
	QuoteService service.QuoteService
}

func NewQuoteController(quoteService service.QuoteService) QuoteController {
	return &QuoteControllerImpl{
		QuoteService: quoteService,
	}
}

func (controller *QuoteControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quote web.QuoteCreateRequest
	err := c.BodyParser(&quote)
	helper.PanicIfError(err)

	ctx := context.Background()
	quoteResponse := controller.QuoteService.Create(ctx, quote)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindById(c *fiber.Ctx) error {
	var quote web.QuoteFindByIdRequest
	id := c.Params("quoteId")
	quote.Id = id

	ctx := context.Background()
	quoteResponse := controller.QuoteService.FindById(ctx, quote)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindAll(c *fiber.Ctx) error {
	ctx := context.Background()
	quoteResponses := controller.QuoteService.FindAll(ctx)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteControllerImpl) FindRandom(c *fiber.Ctx) error {
	ctx := context.Background()
	quoteResponse := controller.QuoteService.FindRandom(ctx)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}