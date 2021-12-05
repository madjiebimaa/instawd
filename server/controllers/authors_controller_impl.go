package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/web"
	"github.com/madjiebimaa/go-random-quotes/services"
)

type AuthorsControllerImpl struct {
	AuthorsService services.AuthorsService
}

func NewAuthorsController(authorsService services.AuthorsService) AuthorsController {
	return &AuthorsControllerImpl{
		AuthorsService: authorsService,
	}
}

func (controller *AuthorsControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var author web.AuthorCreateRequest
	err := c.BodyParser(&author)
	helpers.PanicIfError(err)

	ctx := context.Background()

	authorResponse := controller.AuthorsService.Create(ctx, author)
	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorsControllerImpl) FindById(c *fiber.Ctx) error {
	var author web.AuthorFindByIdRequest
	id := c.Params("authorId")
	author.Id = id

	ctx := context.Background()

	authorResponse := controller.AuthorsService.FindById(ctx, author)
	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorsControllerImpl) FindAll(c *fiber.Ctx) error {
	ctx := context.Background()

	authorResponses := controller.AuthorsService.FindAll(ctx)

	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", authorResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorsControllerImpl) FindBySlug(c *fiber.Ctx) error {
	var author web.AuthorFindBySlugRequest
	slug := c.Params("authorSlug")
	author.Slug = slug

	ctx := context.Background()

	authorResponse := controller.AuthorsService.FindBySlug(ctx, author)
	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorsControllerImpl) FindAuthorAndQuotes(c *fiber.Ctx) error {
	var author web.AuthorFindByIdRequest
	id := c.Params("authorId")
	author.Id = id

	ctx := context.Background()

	authorResponse := controller.AuthorsService.FindAuthorAndQuotes(ctx, author)
	webResponse := helpers.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}
