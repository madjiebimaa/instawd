package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/controllers"
	"github.com/madjiebimaa/go-random-quotes/repositories"
	"github.com/madjiebimaa/go-random-quotes/services"
)

func NewQuotesRoute(app *fiber.App, db *sql.DB, validate *validator.Validate) {
	quoteRepository := repositories.NewQuotesRepository()
	quoteService := services.NewQuotesService(quoteRepository, db, validate)
	quoteController := controllers.NewQuoteController(quoteService)

	app.Post("/api/quotes", quoteController.Create)
	app.Get("/api/quotes", quoteController.FindAll)
	app.Get("/api/quotes/random-quote", quoteController.FindRandom)
	app.Get("/api/quotes/random-quote/author", quoteController.FindRandomAndAuthor)
	app.Get("/api/quotes/:quoteId", quoteController.FindById)
	app.Get("/api/quotes/:quoteId/author", quoteController.FindQuoteAndAuthor)
}
