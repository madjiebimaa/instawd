package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/go-random-quotes/controllers"
	"github.com/madjiebimaa/go-random-quotes/repositories"
	"github.com/madjiebimaa/go-random-quotes/services"
)

func NewAuthorsRoute(app *fiber.App, db *sql.DB, validate *validator.Validate) {
	authorRepository := repositories.NewAuthorsRepository()
	authorService := services.NewAuthorsService(authorRepository, db, validate)
	authorController := controllers.NewAuthorsController(authorService)

	app.Post("/api/authors", authorController.Create)
	app.Get("/api/authors", authorController.FindAll)
	app.Get("/api/authors/:authorId", authorController.FindById)
	app.Get("/api/authors/:authorId/quotes", authorController.FindAuthorAndQuotes)
	app.Get("/api/authors/slug/:authorSlug", authorController.FindBySlug)
}
