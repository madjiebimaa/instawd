package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/instawd/controllers"
	"github.com/madjiebimaa/instawd/repositories"
	"github.com/madjiebimaa/instawd/services"
)

func NewQuotesRoute(app *fiber.App, db *sql.DB, validate *validator.Validate) {
	quoteRepository := repositories.NewQuotesRepository()
	quoteService := services.NewQuotesService(quoteRepository, db)
	quoteController := controllers.NewQuoteController(quoteService, validate)

	app.Post("/api/quotes", quoteController.Create)
	app.Get("/api/quotes", quoteController.FindAll)
	app.Get("/api/quotes/random", quoteController.FindRandom)
	// app.Get("/api/quotes/random/author", quoteController.FindRandomAndAuthor)
	app.Get("/api/quotes/:quoteId", quoteController.FindById)
	// app.Get("/api/quotes/:quoteId/author", quoteController.FindQuoteAndAuthor)
}
