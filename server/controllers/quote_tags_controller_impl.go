package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/web"
	"github.com/madjiebimaa/go-random-quotes/services"
)

type QuoteTagsControllerImpl struct {
	QuoteTagsService services.QuoteTagsService
}

func NewQuoteTagController(quoteTagsService services.QuoteTagsService) QuoteTagsController {
	return &QuoteTagsControllerImpl{
		QuoteTagsService: quoteTagsService,
	}
}

func (controller *QuoteTagsControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quoteTag web.QuoteTagCreateRequest
	err := c.BodyParser(&quoteTag)
	helpers.PanicIfError(err)

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagsService.Create(ctx, quoteTag)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagsControllerImpl) Delete(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var quoteTag web.QuoteTagDeleteRequest
	err := c.BodyParser(&quoteTag)
	helpers.PanicIfError(err)

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagsService.Delete(ctx, quoteTag)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagsControllerImpl) FindById(c *fiber.Ctx) error {
	var quoteTag web.QuoteTagFindByIdRequest
	id := c.Params("quoteTagId")
	quoteTag.Id = id

	ctx := context.Background()

	quoteTagResponse := controller.QuoteTagsService.FindById(ctx, quoteTag)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *QuoteTagsControllerImpl) FindAll(c *fiber.Ctx) error {
	ctx := context.Background()

	quoteTagResponses := controller.QuoteTagsService.FindAll(ctx)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", quoteTagResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}
