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

	// OPENING OF ALL PROGRAMMERS ===============

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// QUOTES ===================================

	app.Post("/api/quotes", quoteController.Create)
	app.Get("/api/quotes", quoteController.FindAll)
	app.Get("/api/quotes/random-quote", quoteController.FindRandom)
	app.Get("/api/quotes/random-quote/author", quoteController.FindRandomAndAuthor)
	app.Get("/api/quotes/:quoteId", quoteController.FindById)
	app.Get("/api/quotes/:quoteId/author", quoteController.FindQuoteAndAuthor)

	// AUTHORS ==================================

	app.Post("/api/authors", authorController.Create)
	app.Get("/api/authors", authorController.FindAll)
	app.Get("/api/authors/:authorId", authorController.FindById)
	app.Get("/api/authors/:authorId/quotes", authorController.FindAuthorAndQuotes)
	app.Get("/api/authors/slug/:authorSlug", authorController.FindBySlug)

	app.Listen(":3000")

	// random with author and author with quotes
}
