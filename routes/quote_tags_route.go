package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/instawd/controllers"
	"github.com/madjiebimaa/instawd/repositories"
	"github.com/madjiebimaa/instawd/services"
)

func NewQuoteTagsRoute(app *fiber.App, db *sql.DB, validate *validator.Validate) {
	quoteTagsRepository := repositories.NewQuoteTagsRepository()
	quoteTagService := services.NewQuoteTagsService(quoteTagsRepository, db)
	quoteTagController := controllers.NewQuoteTagController(quoteTagService, validate)

	app.Post("/api/quote-tags", quoteTagController.Create)
	app.Get("/api/quote-tags", quoteTagController.FindAll)
}
