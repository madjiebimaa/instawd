package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/web"
	"github.com/madjiebimaa/go-random-quotes/service"
)

type AuthorControllerImpl struct {
	AuthorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		AuthorService: authorService,
	}
}

func (controller *AuthorControllerImpl) Create(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)

	var author web.AuthorCreateRequest
	err := c.BodyParser(&author)
	helper.PanicIfError(err)

	ctx := context.Background()

	authorResponse := controller.AuthorService.Create(ctx, author)
	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorControllerImpl) FindById(c *fiber.Ctx) error {
	var author web.AuthorFindByIdRequest
	id := c.Params("authorId")
	author.Id = id

	ctx := context.Background()

	authorResponse := controller.AuthorService.FindById(ctx, author)
	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", authorResponse)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}

func (controller *AuthorControllerImpl) FindAll(c *fiber.Ctx) error {
	ctx := context.Background()

	authorResponses := controller.AuthorService.FindAll(ctx)
	webResponse := helper.ToNewWebResponse(fiber.StatusOK, "OK", authorResponses)

	c.Status(fiber.StatusOK)
	c.Type(fiber.MIMEApplicationJSON)
	return c.JSON(webResponse)
}
