package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/web"
	"github.com/madjiebimaa/go-random-quotes/service"
)

type QuoteTagControllerImpl struct {
	QuoteTagService service.QuoteTagService
}

func NewQuoteTagController(quoteTagService service.QuoteTagService) QuoteTagController {
	return &QuoteTagControllerImpl{
		QuoteTagService: quoteTagService,
	}
}

func (controller *QuoteTagControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quoteTag web.QuoteTagCreateRequest
	err := c.BodyParser(&quoteTag)
	helper.PanicIfError(err)

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagService.Create(ctx, quoteTag)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagControllerImpl) Delete(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quoteTag web.QuoteTagDeleteRequest
	err := c.BodyParser(&quoteTag)
	helper.PanicIfError(err)

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagService.Delete(ctx, quoteTag)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagControllerImpl) FindById(c *fiber.Ctx) error {
	var quoteTag web.QuoteTagFindByIdRequest
	id := c.Params("quoteTagId")
	quoteTag.Id = id

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagService.FindById(ctx, quoteTag)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagControllerImpl) FindAll(c *fiber.Ctx) error {
	ctx := context.Background()

	quoteTagResponses := controller.QuoteTagService.FindAll(ctx)

	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}
