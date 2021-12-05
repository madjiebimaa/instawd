package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/controllers"
	"github.com/madjiebimaa/go-random-quotes/repositories"
	"github.com/madjiebimaa/go-random-quotes/services"
)

func NewQuoteTagsRoute(app *fiber.App, db *sql.DB, validate *validator.Validate) {
	quoteTagsRepository := repositories.NewQuoteTagsRepository()
	quoteTagService := services.NewQuoteTagsService(quoteTagsRepository, db, validate)
	quoteTagController := controllers.NewQuoteTagController(quoteTagService)

	app.Post("/api/quote-tags", quoteTagController.Create)
	app.Get("/api/quote-tags", quoteTagController.FindAll)
	app.Get("/api/quote-tags/:quoteTagId", quoteTagController.FindById)
	app.Delete("/api/quote-tags/", quoteTagController.Delete)
}
