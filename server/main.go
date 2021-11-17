package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	appl "github.com/madjiebimaa/go-random-quotes/app"
	"github.com/madjiebimaa/go-random-quotes/controller"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/repository"
	"github.com/madjiebimaa/go-random-quotes/service"
)

type Response struct {
	Quote  domain.Quote
	Author domain.Author
}

func main() {
	app := fiber.New()
	db := appl.NewDB()
	validate := validator.New()

	quoteRepository := repository.NewQuoteRepository()
	quoteService := service.NewQuoteService(quoteRepository, db, validate)
	quoteController := controller.NewQuoteController(quoteService)

	authorRepository := repository.NewAuthorRepository()
	authorService := service.NewAuthorService(authorRepository, db, validate)
	authorController := controller.NewAuthorController(authorService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/quotes", quoteController.Create)
	app.Get("/api/quotes", quoteController.FindAll)
	app.Get("/api/quotes/:quoteId", quoteController.FindById)
	app.Get("/api/random", quoteController.FindRandom)

	app.Post("/api/authors", authorController.Create)
	app.Get("/api/authors", authorController.FindAll)
	app.Get("/api/authors/:authorId", authorController.FindById)

	app.Listen(":3000")
}
